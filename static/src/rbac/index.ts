import {App, Plugin, reactive} from 'vue'

export class Rbac {
    public permissions = ({} as { [key: string]: boolean })

    public setPermissions(permissions: string[]): void {
        // permissions.forEach(p => {
        //     this.permissions[p] = true
        // })
        this.permissions = permissions.reduce((res, p) => {
            res[p] = true
            return res
        }, {} as { [key: string]: boolean })
    }

    public hasPermission(p: string): boolean {
        return this.permissions[p]
    }
}

const RbacPlugin: Plugin = {
    install(app: App) {
        app.config.globalProperties.$rbac = reactive(new Rbac())
    }
}

export default RbacPlugin