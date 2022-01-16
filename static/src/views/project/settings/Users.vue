<template>
  <h3>
    Users
    <MDBBtn outline="primary" rounded aria-controls="addModal" @click="showAddModal">
      <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
      Add an user
    </MDBBtn>
  </h3>

  <BAddModal caption="Adding the project user" :callback="add" v-model="showAdd">
    <MDBInput label="E-Mail" required v-model="email"/>
  </BAddModal>

  <MDBTable striped sm class="mt-3">
    <thead>
    <tr>
      <th>User</th>
      <th>&ensp;</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="user in options.users" :key="user.email">
      <td>
        <BUser :name="user.fullname" :email="user.email" :avatar-url="user.avatar_url" avatar-size="24"/>
      </td>
      <td class="text-end">
        <a class="action-link link-danger" @click.prevent="showDeleteDialog(user.email)">
          <MDBIcon icon="trash" iconStyle="fas"/>
        </a>
      </td>
    </tr>
    </tbody>
  </MDBTable>

  <BConfirmModal v-model="deleteConfirmShow" yes-btn="Yes" no-btn="Cancel" :callback="deleteUser">
    <p class="font-weight-bold">Are you sure to delete the user "{{ deleteUserCur }}"?</p>
  </BConfirmModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'
import {MDBBtn, MDBIcon, MDBInput, MDBTable} from "mdb-vue-ui-kit";
import api, {ProjectsOptionsOptionsV1} from "@/api";
import BUser from "@/components/BUser.vue";
import BAddModal from "@/components/BAddModal.vue";
import BConfirmModal from "@/components/BConfirmModal.vue";

export default defineComponent({
  name: 'ProjectSettingsUsers',
  components: {BAddModal, BConfirmModal, BUser, MDBBtn, MDBIcon, MDBInput, MDBTable},
  props: {
    options: {
      type: Object as PropType<ProjectsOptionsOptionsV1>,
      required: true
    },
    refreshCb: {
      type: Function as PropType<() => void>,
      required: true
    }
  },
  data() {
    return {
      opts: Object.assign({}, this.options),
      email: '',
      showAdd: false,
      deleteConfirmShow: false,
      deleteUserCur: ''
    }
  },
  methods: {
    showAddModal() {
      this.email = ''
      this.showAdd = true
    },

    // eslint-disable-next-line
    add(): Promise<any> {
      return api.ProjectsUsersAddV1({
        project_id: this.$route.params["id"] as string,
        e_mail: this.email
      }).then(() => {
        this.refreshCb()
      })
    },

    showDeleteDialog(user: string) {
      this.deleteUserCur = user
      this.deleteConfirmShow = true
    },

    // eslint-disable-next-line
    deleteUser(): Promise<any> {
      return api.ProjectsUsersDeleteV1({
        project_id: this.$route.params["id"] as string,
        e_mail: this.deleteUserCur
      }).then(() => {
        this.refreshCb()
      })
    },

  }
});
</script>

<style lang="scss">
.action-link {
  visibility: hidden;
  cursor: pointer;
}

tr:hover a.action-link {
  visibility: visible;
}
</style>