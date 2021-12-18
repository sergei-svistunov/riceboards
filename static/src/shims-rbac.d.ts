import { Rbac } from '@/rbac';

declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        $rbac: Rbac;
    }
}