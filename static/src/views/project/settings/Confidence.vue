<template>
  <h3>
    Confidence levels
    <MDBBtn outline="primary" rounded aria-controls="addModal" @click="showAddModal">
      <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
      Add a level
    </MDBBtn>
  </h3>

  <BAddModal caption="Adding the project confidence level" :callback="add" v-model="showAdd">
    <MDBInput label="Caption" maxLength="255" required v-model="caption"/>
    <MDBInput type="number" label="Weight" class="mt-3 active" required v-model.number="weight"
              formText="Usually a number from 1 to 10"/>
  </BAddModal>

  <MDBTable striped sm class="mt-3">
    <thead>
    <tr>
      <th>Weight</th>
      <th>Caption</th>
      <th>&ensp;</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="level in opts.confident_levels" :key="level.id">
      <td>{{ $filters.formatFloat(level.weight) }}</td>
      <td>{{ level.caption }}</td>
      <td class="text-end">
        <a class="action-link link-warning" @click.prevent="showEditModal(level)">
          <MDBIcon icon="edit" iconStyle="far"/>
        </a>

        <a class="action-link link-danger ms-2" @click.prevent="showDeleteDialog(level)">
          <MDBIcon icon="trash" iconStyle="fas"/>
        </a>
      </td>

    </tr>
    </tbody>
  </MDBTable>

  <BAddModal caption="Editing the project confidence level" :callback="edit" v-model="showEdit" add-btn="Save">
    <MDBInput label="Caption" maxLength="255" required v-model="caption"/>
    <MDBInput type="number" label="Weight" class="mt-3 active" required v-model.number="weight"
              formText="Usually a number from 1 to 10"/>
  </BAddModal>

  <BConfirmModal v-model="deleteConfirmShow" yes-btn="Yes" no-btn="Cancel" :callback="deleteItem">
    <p class="font-weight-bold">Are you sure to delete the level "{{ deleteCur.caption }}"?</p>
  </BConfirmModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'
import {MDBBtn, MDBIcon, MDBInput, MDBTable} from "mdb-vue-ui-kit";
import api, {ProjectsOptionsConfidentV1, ProjectsOptionsOptionsV1} from "@/api";
import BAddModal from "@/components/BAddModal.vue";
import BConfirmModal from "@/components/BConfirmModal.vue";

export default defineComponent({
  name: 'ProjectSettingsConfidence',
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
      weight: 0,

      showAdd: false,
      showEdit: false,
      deleteConfirmShow: false,
      deleteCur: {} as ProjectsOptionsConfidentV1,
    }
  },
  methods: {
    showAddModal() {
      this.caption = ''
      this.showAdd = true
      this.weight = 1
    },

    // eslint-disable-next-line
    add(): Promise<any> {
      return api.ProjectsConfidenceAddV1({
        project_id: this.$route.params["id"] as string,
        caption: this.caption,
        weight: this.weight
      }).then(() => {
        this.refreshCb()
      })
    },

    showEditModal(level: ProjectsOptionsConfidentV1) {
      this.id = level.id
      this.caption = level.caption
      this.weight = level.weight
      this.showEdit = true
    },

    // eslint-disable-next-line
    edit(): Promise<any> {
      return api.ProjectsConfidenceEditV1({
        id: this.id,
        project_id: this.$route.params["id"] as string,
        caption: this.caption,
        weight: this.weight,
      }).then(() => {
        this.refreshCb()
      })
    },

    showDeleteDialog(level: ProjectsOptionsConfidentV1) {
      this.deleteCur = level
      this.deleteConfirmShow = true
    },

    // eslint-disable-next-line
    deleteItem(): Promise<any> {
      return api.ProjectsConfidenceDeleteV1({
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