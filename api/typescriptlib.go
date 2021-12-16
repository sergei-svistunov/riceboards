package api

import (
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/go-qbit/rpc"
)

func typeScriptLibHandler(rpc *rpc.Rpc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, tsLibHeader)

		types := map[string]reflect.Type{}

		for _, path := range rpc.GetPaths() {
			m := rpc.GetMethod(path)

			methodNameParts := strings.Split(strings.TrimPrefix(path, "/"), "/")
			for i, s := range methodNameParts {
				methodNameParts[i] = strings.ToUpper(s[:1]) + s[1:]
			}
			methodName := strings.Join(methodNameParts, "")

			_, _ = io.WriteString(w, "\n\n    // ")
			_, _ = io.WriteString(w, m.Method.Description(r.Context()))
			_, _ = io.WriteString(w, "\n")
			_, _ = io.WriteString(w, `    public static `)
			_, _ = io.WriteString(w, methodName)
			_, _ = io.WriteString(w, "(request: ")
			_, _ = io.WriteString(w, toTsTypeName(m.Request))
			_, _ = io.WriteString(w, "): Promise<")
			_, _ = io.WriteString(w, toTsTypeName(m.Response))
			_, _ = io.WriteString(w, "> {\n        return this.post('")
			_, _ = io.WriteString(w, path)
			_, _ = io.WriteString(w, "', request)\n    }")

			addTsStructTypes(m.Request, types)
			addTsStructTypes(m.Response, types)
		}

		_, _ = io.WriteString(w, "\n}")

		typesNames := make([]string, 0, len(types))
		for t := range types {
			typesNames = append(typesNames, t)
		}
		sort.Strings(typesNames)

		for _, name := range typesNames {
			_, _ = io.WriteString(w, "\n\n")
			if types[name].NumField() == 0 {
				_, _ = io.WriteString(w, "// eslint-disable-next-line\n")
			}
			_, _ = io.WriteString(w, "export interface ")
			_, _ = io.WriteString(w, name)
			_, _ = io.WriteString(w, " {")

			for i := 0; i < types[name].NumField(); i++ {
				field := types[name].Field(i)
				name := field.Tag.Get("json")
				name = strings.Split(name, ",")[0]
				if name == "" {
					name = field.Name
				}
				if name == "-" {
					continue
				}

				_, _ = io.WriteString(w, "\n    ")
				_, _ = io.WriteString(w, name)
				if field.Type.Kind() == reflect.Ptr {
					_, _ = io.WriteString(w, "?")
				}
				_, _ = io.WriteString(w, ": ")
				_, _ = io.WriteString(w, toTsTypeName(field.Type))
				_, _ = io.WriteString(w, ";")

				if description := field.Tag.Get("desc"); description != "" {
					_, _ = io.WriteString(w, "  // ")
					_, _ = io.WriteString(w, description)
				}
			}

			_, _ = io.WriteString(w, "\n}")
		}
	}
}

func toTsTypeName(varType reflect.Type) string {
	if override := typesOverrides[varType.PkgPath()+"."+varType.Name()]; override != "" {
		return override
	}

	typeParts := strings.Split(strings.TrimPrefix(varType.PkgPath(), "riceboards/api/method/"), "/")
	for i, part := range typeParts {
		typeParts[i] = strings.Title(part)
	}
	typePrefix := strings.Join(typeParts, "")

	switch varType.Kind() {
	case reflect.Slice:
		return toTsTypeName(varType.Elem()) + "[]"
	case reflect.Struct:
		sName := varType.Name()
		if sName == "" {
			sName = "Struct_" + strconv.FormatUint(uint64(crc32.ChecksumIEEE([]byte(varType.String()))), 16)
		}
		return typePrefix + strings.Title(sName)
	case reflect.Map:
		return "Record<" + toTsTypeName(varType.Key()) + ", " + toTsTypeName(varType.Elem()) + ">"
	case reflect.Ptr:
		return toTsTypeName(varType.Elem())
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "number"
	case reflect.Interface:
		return "any"
	default:
		panic(fmt.Sprintf("Unknown kind %s", varType.Kind().String()))
	}
}

func addTsStructTypes(st reflect.Type, m map[string]reflect.Type) {
	if typesOverrides[st.PkgPath()+"."+st.Name()] != "" {
		return
	}

	switch st.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Array:
		addTsStructTypes(st.Elem(), m)

	case reflect.Struct:
		m[toTsTypeName(st)] = st
		for i := 0; i < st.NumField(); i++ {
			field := st.Field(i)
			if field.Tag.Get("json") != "-" {
				addTsStructTypes(field.Type, m)
			}
		}

	case reflect.Map:
		addTsStructTypes(st.Elem(), m)

	case reflect.String, reflect.Bool, reflect.Float32, reflect.Float64, reflect.Interface,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return

	default:
		panic(fmt.Sprintf("Unknown kind %s", st.Kind().String()))
	}
}

var (
	typesOverrides = map[string]string{
		"time.Time": "string",
	}

	tsLibHeader = `export class ApiError extends Error {
    private readonly _code: string
    private readonly _message: string
    // eslint-disable-next-line
    private readonly _data: any

    // eslint-disable-next-line
    constructor(code: string, message: string, data: any) {
        super(message);
        this._code = code
        this._message = message
        this._data = data
    }

    get code(): string {
        return this._code;
    }

    get message(): string {
        return this._message;
    }

    get data(): string {
        return this._data;
    }
}

export default class API {
    private static readonly url = "/api"

    // eslint-disable-next-line
    private static post(method: string, request: any): Promise<any> {
        return fetch(
            this.url + method,
            {
                method: 'post',
                headers: {
                    'Content-Type': 'application/json',
                    'X-API-Key': localStorage.getItem('token') || ''
                },
                body: JSON.stringify(request)
            }
        )
            .then(response => {
                // eslint-disable-next-line
                return new Promise<any>((resolve, reject) => {
                    switch (response.status) {
                        case 200:
                            resolve(response)
                            break
                        case 400:
                            response.json().then(err => {
                                reject(new ApiError(err['code'], err['message'], err['data']))
                            })
                            break
                        default:
                            response.text().then(text => {
                                reject(new Error(text || response.statusText))
                            })
                    }
                })
            })
            .then((response) => response.json())
    }`
)
