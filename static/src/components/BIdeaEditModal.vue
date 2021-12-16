<template>
  <MDBModal :id="id" tabindex="-1" :labelledby="`${id}Label`" v-model="show">
    <form @submit.prevent="save" @keypress="formKeyPress">
      <MDBModalHeader>
        <MDBModalTitle :id="`${id}Label`">{{ caption }}</MDBModalTitle>
      </MDBModalHeader>

      <MDBModalBody>
        <slot/>
        <div class="alert alert-danger mt-3" v-if="error">{{ error }}</div>
        <div class="mask text-center" style="background-color: rgba(0, 0, 0, 0.6)" v-if="saving">
          <div class="d-flex justify-content-center align-items-center h-100">
            <MDBSpinner color="light" style="width: 3rem; height: 3rem"/>
          </div>
        </div>
      </MDBModalBody>

      <MDBModalFooter>
        <MDBBtn color="primary" type="submit" :disabled="saving">Save</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {
  MDBBtn,
  MDBModal,
  MDBModalBody,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle,
  MDBSpinner
} from "mdb-vue-ui-kit";
import api, {IdeasEditReqV1} from "@/api";

export default defineComponent({
  name: 'BIdeaEditModal',
  components: {
    MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle, MDBSpinner
  },
  props: {
    id: String,
    caption: String,
    modelValue: Boolean,
    data: {
      type: Object as PropType<IdeasEditReqV1>,
      required: true
    },
    onSave: {
      type: Function as PropType<() => void>,
      required: true,
    },
  },
  emits: ['update:modelValue'],
  data() {
    return {
      show: this.modelValue,
      saving: false,
      error: '',
    }
  },
  watch: {
    show(value) {
      this.$emit('update:modelValue', value)
    },
    modelValue(value) {
      this.show = value
    }
  },
  methods: {
    save() {
      this.saving = true
      this.error = ''
      api.IdeasEditV1(this.data).then(() => {
        this.onSave()
        this.show = false
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.saving = false
      })
    },

    formKeyPress(e: KeyboardEvent) {
      if (e.keyCode === 13 && (e.target as HTMLElement).nodeName != "TEXTAREA") {
        e.preventDefault();
        this.save()
      }
    }
  }
});
</script>