<template>
  <MDBContainer class="mt-3">
    <BContent :loading="loading" :error="error">
      <MDBCol>
        <h6>
        <router-link :to="`/projects/${$route.params['id']}`" class="float-end">
          <MDBIcon icon="arrow-circle-left" iconStyle="fas" class="me-1"/>Back to ideas
        </router-link>
        </h6>

        <h1>{{ options.caption }}</h1>

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
              <ProjectSettingsProject :options="options" :refresh-cb="load"/>
            </MDBTabPane>

            <MDBTabPane tabId="users">
              <ProjectSettingsUsers :options="options" :refresh-cb="load"/>
            </MDBTabPane>

            <MDBTabPane tabId="goals">
              <ProjectSettingsGoals :options="options" :refresh-cb="load"/>
            </MDBTabPane>

            <MDBTabPane tabId="teams">
              <ProjectSettingsTeams :options="options" :refresh-cb="load"/>
            </MDBTabPane>

            <MDBTabPane tabId="confidence">
              <ProjectSettingsConfidence :options="options" :refresh-cb="load"/>
            </MDBTabPane>
          </MDBTabContent>
          <!-- Tabs content -->
        </MDBTabs>
      </MDBCol>
    </BContent>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {MDBCol, MDBContainer, MDBIcon, MDBTabContent, MDBTabItem, MDBTabNav, MDBTabPane, MDBTabs} from "mdb-vue-ui-kit";
import BContent from "@/components/BContent.vue";
import api, {ProjectsOptionsOptionsV1} from "@/api";
import ProjectSettingsProject from "@/views/project/settings/Project.vue";
import {formats} from "@/consts";
import ProjectSettingsUsers from "@/views/project/settings/Users.vue";
import ProjectSettingsGoals from "@/views/project/settings/Goals.vue";
import ProjectSettingsTeams from "@/views/project/settings/Teams.vue";
import ProjectSettingsConfidence from "@/views/project/settings/Confidence.vue";

export default defineComponent({
  name: 'ProjectSettings',
  components: {
    ProjectSettingsConfidence,
    ProjectSettingsTeams,
    ProjectSettingsGoals,
    ProjectSettingsUsers,
    ProjectSettingsProject,
    BContent,
    MDBCol,
    MDBContainer,
    MDBIcon,
    MDBTabContent,
    MDBTabItem,
    MDBTabNav,
    MDBTabPane,
    MDBTabs
  },
  data() {
    return {
      activeTab: 'settings',
      options: {} as ProjectsOptionsOptionsV1,
      loading: true,
      error: '',

    }
  },

  computed: {
    formatsMap(): { [key: number]: string } {
      return formats.reduce((res, f) => {
        res[f.id] = f.label
        return res
      }, {} as { [key: number]: string })
    }
  },

  mounted() {
    this.load()
  },

  methods: {
    load() {
      this.loading = true
      this.error = ''

      api.ProjectsOptionsV1({
        project_id: this.$route.params['id'] as string,
        with_users: true
      }).then(options => {
        this.options = options
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.loading = false
      })
    }
  }
});
</script>
