<template>
  <MDBContainer>
    <BContent :loading="loading" :error="error">
      <MDBCol>
        <MDBTabs v-model="activeTab" vertical>
          <!-- Tabs navs -->
          <MDBTabNav tabsClasses="mb-3" col="3">
            <MDBTabItem tabId="settings" href="settings">Settings</MDBTabItem>
            <MDBTabItem tabId="users" href="users">Users</MDBTabItem>
            <MDBTabItem tabId="goals" href="goals">Goals</MDBTabItem>
            <MDBTabItem tabId="teams" href="teams">Teams</MDBTabItem>
            <MDBTabItem tabId="confidence" href="confidence">Confidence</MDBTabItem>
          </MDBTabNav>
          <!-- Tabs navs -->
          <!-- Tabs content -->
          <MDBTabContent col="9">
            <MDBTabPane tabId="settings">
              <form @submit.prevent="editSettings">
                <MDBInput label="Caption" maxLength="255" required v-model="options.caption"
                          :disabled="editingSettings"/>
                <div class="form-outline mt-3">
                  <select class="form-control active" style="border: #bdbdbd solid thin" id="reach_select"
                          v-model.number="options.reach_format" :disabled="editingSettings">
                    <option value="0">Number</option>
                    <option value="1">Percents</option>
                    <option value="2">Money</option>
                  </select>
                  <label class="form-label bg-white ps-1 pe-1" for="reach_select">Reach format</label>
                </div>

                <MDBInput label="Reach description" maxLength="255" class="mt-3"
                          v-model="options.reach_description" :disabled="editingSettings"/>
                <MDBInput label="Effort description" maxLength="255" class="mt-3"
                          v-model="options.effort_description" :disabled="editingSettings"/>
                <MDBInput label="Money symbol" maxLength="7" required class="mt-3"
                          v-model="options.money_symbol"/>

                <div class="alert alert-danger mt-3" v-if="settingsError">{{ settingsError }}</div>

                <MDBBtn color="primary" type="submit" class="mt-3" :disabled="editingSettings">Save</MDBBtn>
              </form>
            </MDBTabPane>

            <MDBTabPane tabId="users">Coming soon</MDBTabPane>

            <MDBTabPane tabId="goals">Coming soon</MDBTabPane>

            <MDBTabPane tabId="teams">Coming soon</MDBTabPane>

            <MDBTabPane tabId="confidence">Coming soon</MDBTabPane>
          </MDBTabContent>
          <!-- Tabs content -->
        </MDBTabs>
      </MDBCol>
    </BContent>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {
  MDBBtn,
  MDBCol,
  MDBContainer,
  MDBInput,
  MDBTabContent,
  MDBTabItem,
  MDBTabNav,
  MDBTabPane,
  MDBTabs
} from "mdb-vue-ui-kit";
import BContent from "@/components/BContent.vue";
import api, {IdeasOptionsOptionsV1} from "@/api";

export default defineComponent({
  name: 'ProjectSettings',
  components: {
    BContent,
    MDBBtn,
    MDBCol,
    MDBContainer,
    MDBInput,
    MDBTabContent,
    MDBTabItem,
    MDBTabNav,
    MDBTabPane,
    MDBTabs
  },
  data() {
    return {
      activeTab: 'settings',
      options: {} as IdeasOptionsOptionsV1,
      loading: true,
      error: '',

      editingSettings: false,
      settingsError: ''
    }
  },

  mounted() {
    this.load()
  },

  methods: {
    load() {
      this.loading = true
      this.error = ''

      api.IdeasOptionsV1({
        project_id: parseInt(this.$route.params['id'] as string)
      }).then(options => {
        this.options = options
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.loading = false
      })
    },

    editSettings() {
      this.editingSettings = true
      this.settingsError = ''

      api.ProjectsEditV1({
        id: parseInt(this.$route.params['id'] as string),
        caption: this.options.caption,
        effort_description: this.options.effort_description,
        money_symbol: this.options.money_symbol,
        reach_description: this.options.reach_description,
        reach_format: this.options.reach_format
      }).catch(err => {
        this.settingsError = err
      }).finally(() => {
        this.editingSettings = false
      })
    }
  }
});
</script>
