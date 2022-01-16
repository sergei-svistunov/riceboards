<template>
  <MDBContainer>
    <BContent :loading="loading" :error="error">
      <MDBCol md="12" class="mt-3">
        <h1>
          My projects
          <MDBBtn outline="primary" rounded size="md" style="margin-bottom: 0.5em;"
                  aria-controls="addModal" @click="addForm">
            <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
            Add a project
          </MDBBtn>
        </h1>

        <MDBTable striped sm class="mt-3">
          <tbody>
          <tr v-for="project in projects" :key="project.id">
            <td class="position-relative">
              <router-link :to="`/projects/${project.id}`" class="stretched-link">
                <MDBBadge color="primary">{{ project.id }}</MDBBadge>
              </router-link>
            </td>

            <td class="position-relative">
              <router-link :to="`/projects/${project.id}`" class="stretched-link">{{ project.caption }}</router-link>
            </td>

            <td class="text-end">
              <MDBBtn color="light" floating size="sm" title="Settings" tag="a"
                      :href="`/projects/${project.id}/settings`">
                <MDBIcon icon="cog" iconStyle="fas" size="md"/>
              </MDBBtn>
            </td>
          </tr>
          </tbody>
        </MDBTable>

        <div class="alert alert-warning" v-if="projects.length === 0">No projects was found</div>

      </MDBCol>

    </BContent>
  </MDBContainer>

  <MDBModal size="lg" id="addModal" tabindex="-1" labelledby="addModalLabel" v-model="addModal">
    <form @submit.prevent="add" @keypress="formKeyPress">
      <MDBModalHeader>
        <MDBModalTitle id="addModalLabel">Adding a project</MDBModalTitle>
      </MDBModalHeader>

      <MDBModalBody>
        <MDBInput label="Id" required v-model="addData.id" :disabled="adding"/>
        <MDBInput label="Caption" required class="mt-3" v-model="addData.caption" :disabled="adding"/>

        <div class="alert alert-danger mt-3" v-if="addError">{{ addError }}</div>
      </MDBModalBody>

      <MDBModalFooter>
        <MDBBtn color="primary" type="submit" :disabled="adding">Add</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<script lang="ts">
import api, {ProjectsListProjectV1} from '@/api';
import {defineComponent} from 'vue'
import BContent from "@/components/BContent.vue";
import {
  MDBBadge,
  MDBBtn,
  MDBCol,
  MDBContainer,
  MDBIcon,
  MDBInput,
  MDBModal,
  MDBModalBody,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle,
  MDBTable
} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'Projects',
  components: {
    BContent,
    MDBBadge,
    MDBBtn,
    MDBCol,
    MDBContainer,
    MDBIcon,
    MDBInput,
    MDBModal,
    MDBModalBody,
    MDBModalFooter,
    MDBModalHeader,
    MDBModalTitle,
    MDBTable
  },
  data() {
    return {
      loading: true,
      error: '',
      projects: [] as ProjectsListProjectV1[],

      addModal: false,
      addData: {
        id: '',
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
    },

    formKeyPress(e: KeyboardEvent) {
      if (e.keyCode === 13 && (e.target as HTMLElement).nodeName != "TEXTAREA") {
        e.preventDefault();
        this.add()
      }
    }
  }
});
</script>
