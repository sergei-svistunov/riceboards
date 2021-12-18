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
          <MDBDropdownItem href="#logout"
                           @click.prevent="$store.commit('removeUser');$rbac.setPermissions([]);showLogin()">Logout
          </MDBDropdownItem>
        </MDBDropdownMenu>
      </MDBDropdown>

      <MDBNavbarItem v-else>
        <div id="g_id_onload"
             data-client_id="709309428199-nnm1q8kikmg9n95f287suqmvoi0tm2go.apps.googleusercontent.com"
             data-context="signin"
             data-ux_mode="popup"
             data-auto_prompt="false">
        </div>

        <div class="g_id_signin"
             data-type="standard"
             data-shape="pill"
             data-theme="outline"
             data-size="large"
             data-logo_alignment="left">
        </div>
      </MDBNavbarItem>
    </MDBNavbarNav>
  </MDBNavbar>

  <router-view class="mt-3"/>
</template>

<script>
import {defineComponent} from "vue"
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
import googleOneTap from 'google-one-tap'
import api from "@/api";

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
      googleOneTap({
        client_id: '1030387876775-48acocqe8ue99usaja3a70pis4125drl.apps.googleusercontent.com', // required
        auto_select: false, // optional
        cancel_on_tap_outside: false, // optional
        context: 'signin', // optional
      }, (response) => {
        api.AuthGoogleV1({
          credential: response.credential
        }).then(user => {
          this.$store.commit('setUser', user)
          this.$rbac.setPermissions(user.permissions)
        }).catch(err => {
          alert(err)
        })
      });
    }
  }
})
</script>

<style lang="scss">
@import '~@/../mdb/scss/index.free.scss';
</style>
