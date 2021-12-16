<template>
  <MDBContainer>
    <BContent :loading="loading" :error="error">
      <MDBCol md="2" class="sidebar">
        <h3 class="mt-3">{{ project.caption }}</h3>
        <ul class="menu">
          <li>
            <router-link :to="`/projects/${$route.params['id']}`">
              <MDBIcon icon="lightbulb" iconStyle="far" class="me-3"/>
              Ideas
            </router-link>
          </li>
          <li>
            <router-link to="#">
              <MDBIcon icon="cog" iconStyle="fas" class="me-3"/>
              Settings
            </router-link>
            <ul class="submenu">
              <li>
                <router-link to="#">Goals</router-link>
              </li>
              <li>
                <router-link to="#">Teams</router-link>
              </li>
              <li>
                <router-link to="#">Confidence stages</router-link>
              </li>
            </ul>
          </li>
        </ul>
      </MDBCol>

      <MDBCol md="10">
        <router-view/>
      </MDBCol>
    </BContent>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {MDBCol, MDBContainer, MDBIcon} from "mdb-vue-ui-kit";
import api, {ProjectsGetProjectV1} from "@/api";
import BContent from "@/components/BContent.vue";

export default defineComponent({
  name: 'Project',
  components: {BContent, MDBContainer, MDBCol, MDBIcon},
  data() {
    return {
      project: {} as ProjectsGetProjectV1,
      loading: true,
      error: ''
    }
  },
  mounted() {
    api.ProjectsGetV1({id: parseInt(this.$route.params['id'] as string)}).then(project => {
      this.project = project
    }).catch(err => {
      this.error = err
    }).finally(() => {
      this.loading = false
    })
  }
});
</script>

<style lang="scss">
@import "~mdb-vue-ui-kit/src/scss/index.free";

.sidebar {
  height: calc(100vh - 60px);
  background-color: $light;
  margin-top: -15px;
  overflow-y: auto;
  overflow-x: hidden;
}

.sidebar .menu {
  list-style: none;
  position: relative;
  padding: 0 0.2rem;
  margin: 0;
}

.sidebar .menu a {
  display: flex;
  align-items: center;
  cursor: pointer;
  font-size: .89rem;
  padding: 1rem 1.5rem;
  height: 3rem;
  color: unset;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  border-radius: 5px;
  transition: all .3s linear;
}

.sidebar .submenu {
  list-style: none;
  position: relative;
  padding: 0 0.2rem;
  margin: 0;
}

.sidebar .submenu a {
  font-size: .78rem;
  height: 1.5rem;
  padding-left: 3rem
}
</style>