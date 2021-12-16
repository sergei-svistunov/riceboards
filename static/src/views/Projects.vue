<template>
  <MDBContainer>
    <BContent :loading="loading" :error="error">
      <h1>Projects
        <MDBBtn color="primary" floating size="md" style="margin-bottom: -0.5em;" title="Add a new project" aria-controls="addModal"
                @click="addForm">
          <MDBIcon icon="plus" iconStyle="fas"/>
        </MDBBtn>
      </h1>

      <MDBModal size="lg" id="addModal" tabindex="-1" labelledby="addModalLabel" v-model="addModal">
        <form @submit.prevent="add">
          <MDBModalHeader>
            <MDBModalTitle id="addModalLabel">Adding a project</MDBModalTitle>
          </MDBModalHeader>

          <MDBModalBody>
            <MDBInput label="Caption" required v-model="addData.caption" :disabled="adding"/>

            <div class="alert alert-danger mt-3" v-if="addError">{{ addError }}</div>
          </MDBModalBody>

          <MDBModalFooter>
            <MDBBtn color="primary" type="submit" :disabled="adding">Add</MDBBtn>
          </MDBModalFooter>
        </form>
      </MDBModal>

      <MDBCol md="12" class="mt-3" v-for="project in projects" :key="project.id">
        <router-link :to="`/projects/${project.id}`">{{ project.caption }}</router-link>
      </MDBCol>

      <div class="alert alert-warning" v-if="projects.length === 0">No projects was found</div>
    </BContent>
  </MDBContainer>
</template>

<script lang="ts">
import api, {ProjectsListProjectV1} from '@/api';
import {defineComponent} from 'vue'
import BContent from "@/components/BContent.vue";
import {
  MDBBtn,
  MDBCol,
  MDBContainer,
  MDBIcon,
  MDBInput,
  MDBModal,
  MDBModalBody,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle
} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'Projects',
  components: {
    BContent,
    MDBBtn,
    MDBCol,
    MDBContainer,
    MDBIcon,
    MDBInput,
    MDBModal,
    MDBModalBody,
    MDBModalFooter,
    MDBModalHeader,
    MDBModalTitle
  },
  data() {
    return {
      loading: true,
      error: '',
      projects: [] as ProjectsListProjectV1[],

      addModal: false,
      addData: {
        caption: ''
      },
      adding: false,
      addError: '',
    }
  },
  mounted() {
    this.load()
  },
  methods: {
    load() {
      this.loading = true
      this.error = ''

      api.ProjectsListV1({}).then(projects => {
        this.projects = projects
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.loading = false
      })
    },

    addForm() {
      this.addModal = true
    },

    add() {
      this.adding = true
      this.addError = ''

      api.ProjectsAddV1(this.addData).then(() => {
        this.addModal = false
        this.load()
      }).catch(err => {
        this.addError = err
      }).finally(() => {
        this.adding = false
      })
    }
  }
});
</script>
