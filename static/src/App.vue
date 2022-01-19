<template>
  <MDBNavbar expand="lg" light bg="light" container>
    <MDBNavbarBrand href="/">RICE boards</MDBNavbarBrand>

    <MDBNavbarToggler
        @click="collapse = !collapse"
        target="#navbarSupportedContent"
    ></MDBNavbarToggler>

    <MDBCollapse v-model="collapse" id="navbarSupportedContent">
      <MDBNavbarNav class="mb-2 mb-lg-0">
        <MDBNavbarItem to="/projects" :active="$route.path.startsWith('/projects')"
                       v-if="$rbac.hasPermission('project.view')">Projects
        </MDBNavbarItem>
      </MDBNavbarNav>
    </MDBCollapse>

    <MDBNavbarNav right>
      <MDBDropdown class="nav-item" align="end" v-model="profileDropdown" v-if="curUser.id">
        <MDBDropdownToggle tag="a" class="nav-link" @click="profileDropdown = !profileDropdown"
        ><img
            :src="curUser.avatarUrl"
            class="rounded-circle"
            height="24"
            :alt="curUser.name"
            loading="lazy"
        />
        </MDBDropdownToggle>
        <MDBDropdownMenu animation>
          <MDBDropdownItem href="#logout" @click.prevent="doLogout">Logout
          </MDBDropdownItem>
        </MDBDropdownMenu>
      </MDBDropdown>

      <MDBNavbarItem v-else>
        <div id="gsi-btn"/>
      </MDBNavbarItem>
    </MDBNavbarNav>
  </MDBNavbar>

  <router-view/>
</template>

<script>
import {defineComponent} from 'vue'
import {
  MDBCollapse,
  MDBDropdown,
  MDBDropdownItem,
  MDBDropdownMenu,
  MDBDropdownToggle,
  MDBNavbar,
  MDBNavbarBrand,
  MDBNavbarItem,
  MDBNavbarNav,
  MDBNavbarToggler
} from 'mdb-vue-ui-kit'
import api from "@/api";
import google_on_tap from "@/gsi";

export default defineComponent({
  name: 'App',
  components: {
    MDBCollapse,
    MDBDropdown,
    MDBDropdownItem,
    MDBDropdownMenu,
    MDBDropdownToggle,
    MDBNavbar,
    MDBNavbarBrand,
    MDBNavbarItem,
    MDBNavbarNav,
    MDBNavbarToggler
  },
  data() {
    return {
      collapse: false,
      profileDropdown: false
    }
  },
  computed: {
    curUser() {
      return this.$store.state.user
    }
  },
  mounted() {
    if (localStorage.getItem('token')) {
      api.AuthMeV1({}).then(user => {
        this.$store.commit('setUser', user)
        this.$rbac.setPermissions(user.permissions)
      }).catch(() => {
        this.$store.commit('removeUser')
        this.$rbac.setPermissions([])
        this.showLogin()
      })
    } else {
      this.showLogin()
    }
  },
  methods: {
    showLogin() {
      google_on_tap({
        client_id: '1030387876775-48acocqe8ue99usaja3a70pis4125drl.apps.googleusercontent.com',
        callback: this.doLogin
      })
    },

    doLogin(response) {
      api.AuthGoogleV1({
        credential: response.credential
      }).then(user => {
        this.$store.commit('setUser', user)
        this.$rbac.setPermissions(user.permissions)
      }).catch(err => {
        alert(err)
      })
    },

    doLogout() {
      this.$store.commit('removeUser')
      this.$rbac.setPermissions([])
      this.$router.push('/')
    }

  }
})
</script>

<style lang="scss">
@import 'mdb/scss/index.free';
</style>
