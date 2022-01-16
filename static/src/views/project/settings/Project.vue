<template>
  <h3>Settings</h3>
  <form @submit.prevent="editSettings">
    <MDBInput label="Caption" maxLength="255" required class="mt-3" v-model="opts.caption"
              :disabled="editingSettings"/>
    <div class="form-outline mt-3">
      <select class="form-control active" style="border: #bdbdbd solid thin" id="reach_select"
              v-model.number="opts.reach_format" :disabled="editingSettings">
        <option v-for="format in formats" :key="format.id" :value="format.id">{{ format.label }}</option>
      </select>
      <label class="form-label bg-white ps-1 pe-1" for="reach_select">Reach format</label>
    </div>

    <MDBInput label="Reach description" maxLength="255" class="mt-3"
              v-model="opts.reach_description" :disabled="editingSettings"/>
    <MDBInput label="Effort description" maxLength="255" class="mt-3"
              v-model="opts.effort_description" :disabled="editingSettings"/>
    <MDBInput label="Money symbol" maxLength="7" required class="mt-3"
              v-model="opts.money_symbol"/>

    <div class="alert alert-danger mt-3" v-if="settingsError">{{ settingsError }}</div>

    <MDBBtn color="primary" type="submit" class="mt-3" :disabled="editingSettings">Save</MDBBtn>
  </form>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'
import {MDBBtn, MDBInput} from "mdb-vue-ui-kit";
import api, {ProjectsOptionsOptionsV1} from "@/api";
import {formats} from "@/consts";

export default defineComponent({
  name: 'ProjectSettingsProject',
  components: {MDBInput, MDBBtn},
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
      editingSettings: false,
      settingsError: '',

      formats: formats
    }
  },
  methods: {
    editSettings() {
      this.editingSettings = true
      this.settingsError = ''

      api.ProjectsEditV1({
        id: this.$route.params['id'] as string,
        caption: this.opts.caption,
        effort_description: this.opts.effort_description,
        money_symbol: this.opts.money_symbol,
        reach_description: this.opts.reach_description,
        reach_format: this.opts.reach_format
      }).then(() => {
        this.refreshCb()
      }).catch(err => {
        this.settingsError = err
      }).finally(() => {
        this.editingSettings = false
      })
    }
  }
});
</script>
