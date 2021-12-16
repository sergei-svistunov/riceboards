import {createStore} from 'vuex'
import {AuthGoogleUserV1, AuthMeUserV1} from "@/api";

export default createStore({
    state: {
        user: {
            id: 0,
            name: '',
            email: '',
            avatarUrl: ''
        },
        permissions: {} as { [key: string]: boolean }
    },
    mutations: {
        setUser(state, user: AuthGoogleUserV1 | AuthMeUserV1) {
            state.user.id = user.id
            state.user.name = user.fullname
            state.user.email = user.email
            state.user.avatarUrl = user.avatar_url
            if ((<AuthGoogleUserV1>user).token) {
                localStorage.setItem('token', (<AuthGoogleUserV1>user).token)
            }

            // const permissions: { [key: string]: boolean } = {}
            // if (user.permissions)
            //   user.permissions.forEach(v => {
            //     permissions[v] = true
            //   })
            //
            // state.permissions = permissions
        },

        removeUser(state) {
            state.user = {
                id: 0,
                name: "",
                email: '',
                avatarUrl: ""
            }
            localStorage.removeItem("token")
        }
    },
    actions: {},
    modules: {}
})
