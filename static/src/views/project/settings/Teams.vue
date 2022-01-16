<template>
  <h3>
    Teams
    <MDBBtn outline="primary" rounded aria-controls="addModal" @click="showAddModal">
      <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
      Add a team
    </MDBBtn>
  </h3>

  <BAddModal caption="Adding the project team" :callback="add" v-model="showAdd">
    <MDBInput label="Caption" maxLength="255" required v-model="caption"/>
  </BAddModal>

  <MDBTable striped sm class="mt-3">
    <thead>
    <tr>
      <th>Caption</th>
      <th>&ensp;</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="team in opts.teams" :key="team.id">
      <td>{{ team.caption }}</td>
      <td class="text-end">
        <a class="action-link link-warning" @click.prevent="showEditModal(team)">
          <MDBIcon icon="edit" iconStyle="far"/>
        </a>

        <a class="action-link link-danger ms-2" @click.prevent="showDeleteDialog(team)">
          <MDBIcon icon="trash" iconStyle="fas"/>
        </a>
      </td>

    </tr>
    </tbody>
  </MDBTable>

  <BAddModal caption="Editing the project team" :callback="edit" v-model="showEdit" add-btn="Save">
    <MDBInput label="Caption" maxLength="255" required v-model="caption"/>
  </BAddModal>

  <BConfirmModal v-model="deleteConfirmShow" yes-btn="Yes" no-btn="Cancel" :callback="deleteItem">
    <p class="font-weight-bold">Are you sure to delete the goal "{{ deleteCur.caption }}"?</p>
  </BConfirmModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'
import {MDBBtn, MDBIcon, MDBInput, MDBTable} from "mdb-vue-ui-kit";
import api, {ProjectsOptionsOptionsV1, ProjectsOptionsTeamV1} from "@/api";
import BAddModal from "@/components/BAddModal.vue";
import BConfirmModal from "@/components/BConfirmModal.vue";

export default defineComponent({
  name: 'ProjectSettingsTeams',
  components: {BAddModal, BConfirmModal, MDBBtn, MDBIcon, MDBInput, MDBTable},
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

      id: 0,
      caption: '',

      showAdd: false,
      showEdit: false,
      deleteConfirmShow: false,
      deleteCur: {} as ProjectsOptionsTeamV1,
    }
  },
  methods: {
    showAddModal() {
      this.caption = ''
      this.showAdd = true
    },

    // eslint-disable-next-line
    add(): Promise<any> {
      return api.ProjectsTeamsAddV1({
        project_id: this.$route.params["id"] as string,
        caption: this.caption,
      }).then(() => {
        this.refreshCb()
      })
    },

    showEditModal(team: ProjectsOptionsTeamV1) {
      this.id = team.id
      this.caption = team.caption
      this.showEdit = true
    },

    // eslint-disable-next-line
    edit(): Promise<any> {
      return api.ProjectsTeamsEditV1({
        id: this.id,
        project_id: this.$route.params["id"] as string,
        caption: this.caption,
      }).then(() => {
        this.refreshCb()
      })
    },

    showDeleteDialog(team: ProjectsOptionsTeamV1) {
      this.deleteCur = team
      this.deleteConfirmShow = true
    },

    // eslint-disable-next-line
    deleteItem(): Promise<any> {
      return api.ProjectsTeamsDeleteV1({
        project_id: this.$route.params["id"] as string,
        id: this.deleteCur.id
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