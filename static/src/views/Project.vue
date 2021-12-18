<template>
  <MDBContainer class="mt-3">
    <BContent :loading="loading" :error="error">
      <MDBCol md="12">
        <h1>
          {{ project.caption }}
          <MDBBtn color="light" floating size="md" style="margin-bottom: 0.5em;" class="float-end" title="Settings">
            <MDBIcon icon="cog" iconStyle="fas" size="2x"/>
          </MDBBtn>
        </h1>

        <router-view/>
      </MDBCol>
    </BContent>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {MDBBtn, MDBCol, MDBContainer, MDBIcon} from "mdb-vue-ui-kit";
import api, {ProjectsGetProjectV1} from "@/api";
import BContent from "@/components/BContent.vue";

export default defineComponent({
  name: 'Project',
  components: {BContent, MDBBtn, MDBCol, MDBContainer, MDBIcon},
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