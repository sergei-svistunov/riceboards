<template>
  <h3>
    Goals
    <MDBBtn outline="primary" rounded aria-controls="addModal" @click="showAddModal">
      <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
      Add a goal
    </MDBBtn>
  </h3>

  <BAddModal caption="Adding the project goal" :callback="add" v-model="showAdd">
    <MDBInput label="Caption" maxLength="255" required v-model="caption"/>

    <div class="form-outline mt-3">
      <select class="form-control active" style="border: #bdbdbd solid thin" id="confident_select"
              v-model.number="format">
        <option v-for="format in formats" :key="format.id" :value="format.id">
          {{ format.label }}
        </option>
      </select>
      <label class="form-label bg-white ps-1 pe-1" for="confident_select">Format</label>
    </div>

    <MDBInput type="number" label="Divider" class="mt-3 active" required v-model.number="divider"
              formText="The number that coincides 10 points"/>

  </BAddModal>

  <MDBTable striped sm class="mt-3">
    <thead>
    <tr>
      <th>Caption</th>
      <th>Format</th>
      <th>Divider</th>
      <th>&ensp;</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="goal in opts.goals" :key="goal.id">
      <td>{{ goal.caption }}</td>
      <td>{{ formatsMap[goal.format] }}</td>
      <td>{{ $filters.formatFloat(goal.divider) }}</td>
      <td class="text-end">
        <a class="action-link link-warning" @click.prevent="showEditModal(goal)">
          <MDBIcon icon="edit" iconStyle="far"/>
        </a>

        <a class="action-link link-danger ms-2" @click.prevent="showDeleteDialog(goal)">
          <MDBIcon icon="trash" iconStyle="fas"/>
        </a>
      </td>

    </tr>
    </tbody>
  </MDBTable>

  <BAddModal caption="Editing the project goal" :callback="edit" v-model="showEdit" add-btn="Save">
    <MDBInput label="Caption" maxLength="255" required v-model="caption"/>

    <div class="form-outline mt-3">
      <select class="form-control active" style="border: #bdbdbd solid thin" id="goal_format_edit"
              v-model.number="format">
        <option v-for="format in formats" :key="format.id" :value="format.id">
          {{ format.label }}
        </option>
      </select>
      <label class="form-label bg-white ps-1 pe-1" for="goal_format_edit">Format</label>
    </div>

    <MDBInput type="number" label="Divider" class="mt-3 active" required v-model.number="divider"
              formText="The number that coincides 10 points"/>

  </BAddModal>

  <BConfirmModal v-model="deleteConfirmShow" yes-btn="Yes" no-btn="Cancel" :callback="deleteItem">
    <p class="font-weight-bold">Are you sure to delete the goal "{{ deleteCur.caption }}"?</p>
  </BConfirmModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'
import {MDBBtn, MDBIcon, MDBInput, MDBTable} from "mdb-vue-ui-kit";
import api, {ProjectsOptionsGoalV1, ProjectsOptionsOptionsV1} from "@/api";
import BAddModal from "@/components/BAddModal.vue";
import BConfirmModal from "@/components/BConfirmModal.vue";
import {formats} from "@/consts";

export default defineComponent({
  name: 'ProjectSettingsGoals',
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
      format: 0,
      divider: 0,

      showAdd: false,
      showEdit: false,
      deleteConfirmShow: false,
      deleteCur: {} as ProjectsOptionsGoalV1,
      formats: formats,
      formatsMap: formats.reduce((res, f) => {
        res[f.id] = f.label
        return res
      }, {} as { [key: number]: string })
    }
  },
  methods: {
    showAddModal() {
      this.caption = ''
      this.format = 0
      this.showAdd = true
      this.divider = 1
    },

    // eslint-disable-next-line
    add(): Promise<any> {
      return api.ProjectsGoalsAddV1({
        project_id: this.$route.params["id"] as string,
        caption: this.caption,
        format: this.format,
        divider: this.divider
      }).then(() => {
        this.refreshCb()
      })
    },

    showEditModal(goal: ProjectsOptionsGoalV1) {
      this.id = goal.id
      this.caption = goal.caption
      this.format = goal.format
      this.divider = goal.divider
      this.showEdit = true
    },

    // eslint-disable-next-line
    edit(): Promise<any> {
      return api.ProjectsGoalsEditV1({
        id: this.id,
        project_id: this.$route.params["id"] as string,
        caption: this.caption,
        format: this.format,
        divider: this.divider,
      }).then(() => {
        this.refreshCb()
      })
    },

    showDeleteDialog(goal: ProjectsOptionsGoalV1) {
      this.deleteCur = goal
      this.deleteConfirmShow = true
    },

    // eslint-disable-next-line
    deleteItem(): Promise<any> {
      return api.ProjectsGoalsDeleteV1({
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