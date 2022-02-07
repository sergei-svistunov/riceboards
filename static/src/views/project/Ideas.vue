<template>
  <MDBContainer class="mt-3">
    <BContent :loading="loading || optionsLoading" :error="error || optionsError">
      <MDBCol>
        <MDBBtn color="light" floating size="md" style="margin-bottom: 0.5em;" class="float-end" title="Settings"
                tag="a" :href="`/projects/${$route.params['id']}/settings`">
          <MDBIcon icon="cog" iconStyle="fas" size="2x"/>
        </MDBBtn>

        <h1>
          {{ options.caption }}
          <MDBBtn outline="primary" rounded size="md" style="margin-bottom: 0.5em;"
                  aria-controls="addModal" @click="addForm">
            <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
            Add an idea
          </MDBBtn>
        </h1>

        <MDBModal size="lg" id="addModal" tabindex="-1" labelledby="addModalLabel" v-model="addModal">
          <form @submit.prevent="add" @keypress="formKeyPress">
            <MDBModalHeader>
              <MDBModalTitle id="addModalLabel">Adding an idea</MDBModalTitle>
            </MDBModalHeader>

            <MDBModalBody>
              <MDBCol>
                <MDBInput label="Caption" required v-model="addData.caption" :disabled="adding"/>
              </MDBCol>
              <div class="alert alert-danger mt-3" v-if="addError">{{ addError }}</div>
            </MDBModalBody>

            <MDBModalFooter>
              <MDBBtn color="primary" type="submit" :disabled="adding">Add</MDBBtn>
            </MDBModalFooter>
          </form>
        </MDBModal>

        <MDBTable striped sm responsive class="ideas-tbl">
          <thead>
          <tr>
            <th>Priority</th>
            <th>Idea</th>
            <th></th>
            <th class="text-nowrap">
              Reach
              <BComment class="ms-1" :text="options.reach_description" v-if="options.reach_description"/>
            </th>
            <th>Impact</th>
            <th>Confidence</th>
            <th class="text-nowrap">
              Effort
              <BComment class="ms-1" :text="options.effort_description" v-if="options.effort_description"/>
            </th>
            <th>Score</th>
            <th>&ensp;</th>
          </tr>
          </thead>

          <tbody>
          <tr v-for="(idea, i) in ideasView" :key="idea.id" :id="`idea${idea.id}`"
              :class="{
                  'table-warning': (idea.score === undefined || !idea.filled_effort) && !idea.done,
                  'table-success': idea.done,
                  'current': curIdeaId === idea.id
              }">
            <th>{{ i + 1 }}</th>
            <th class="position-relative">
              <BUser :name="idea.owner.fullname" :email="idea.owner.email" :avatar-url="idea.owner.avatar_url"
                     :hide-name="true" avatar-size="20"/>
              {{ idea.caption }}
              <BComment class="ms-1" :text="idea.comment" v-if="idea.comment"/>
              <a class="edit-link" @click.prevent="editIdea(idea)">
                <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
              </a>
            </th>

            <td class="position-relative">
              <a :href="idea.link" v-if="idea.link" target="_blank" class="ms-2">
                <MDBIcon icon="external-link-alt" iconStyle="fas"/>
              </a>
              <MDBIcon icon="dev" iconStyle="fab" v-if="idea.ready_for_dev" title="Ready for development"
                       class="ms-2"/>
              <MDBIcon icon="check-circle" iconStyle="fas" v-if="idea.done" title="Done" class="ms-2"/>
            </td>

            <td class="position-relative">
              <template v-if="idea.reach !== undefined">{{ formatNumber(idea.reach, options.reach_format) }}</template>
              <span class="text-muted" v-else>&mdash;</span>
              <BComment class="ms-1" :text="idea.reach_comment" v-if="idea.reach_comment"/>

              <a class="edit-link" @click.prevent="editReach(idea)">
                <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
              </a>
            </td>

            <td class="position-relative text-nowrap">
              <ul class="list-unstyled mb-0" v-if="filteredGoals(idea).length">
                <li v-for="goal in filteredGoals(idea)" :key="goal.id">
                  <!--                  <BUser :email="idea.goals[goal.id].owner.email" :name="idea.goals[goal.id].owner.fullname"
                                           :avatar-url="idea.goals[goal.id].owner.avatar_url" hide-name avatar-size="20"/>-->
                  <MDBBadge color="success" class="ms-1" style="width: 3em">
                    {{ idea.goals[goal.id].value * 10 / goal.divider }}
                  </MDBBadge>
                  {{ goal.caption }}:&nbsp;{{ formatNumber(idea.goals[goal.id].value, goal.format) }}
                  <BComment class="ms-1" :text="idea.goals[goal.id]?.comment" v-if="idea.goals[goal.id]?.comment"/>
                </li>
              </ul>
              <template v-else>
                &mdash;
              </template>
              <a class="edit-link" @click.prevent="editGoals(idea)">
                <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
              </a>
            </td>

            <td class="position-relative">
              <template v-if="idea.confident !== undefined">
                <MDBBadge color="secondary" pill>{{ confidentLevelsMap[idea.confident].weight }}</MDBBadge>
                {{ confidentLevelsMap[idea.confident].caption }}
              </template>
              <span class="text-muted" v-else>&mdash;</span>
              <BComment class="ms-1" :text="idea.confident_comment" v-if="idea.confident_comment"/>

              <a class="edit-link" @click.prevent="editConfident(idea)">
                <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
              </a>
            </td>

            <td class="position-relative text-nowrap">
              <ul class="list-unstyled mb-0" v-if="filteredTeams(idea).length">
                <li v-for="team in filteredTeams(idea)" :key="team.id">
                  <BUser :email="idea.teams[team.id].owner.email" :name="idea.teams[team.id].owner.fullname"
                         :avatar-url="idea.teams[team.id].owner.avatar_url" hide-name avatar-size="16"
                         v-if="idea.teams[team.id]"/>
                  <span :class="{'ms-2': idea.teams[team.id], 'text-danger': !idea.teams[team.id]}">{{ team.caption }}:&nbsp;{{
                      idea.teams[team.id]?.capacity || '&mdash;'
                    }}</span>
                  <BComment class="ms-1" :text="idea.teams[team.id]?.comment" v-if="idea.teams[team.id]?.comment"/>
                </li>
              </ul>
              <template v-else>
                &mdash;
              </template>
              <a class="edit-link" @click.prevent="editEffort(idea)">
                <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
              </a>
            </td>

            <th>
              <template v-if="idea.score === undefined">
                &mdash;
              </template>
              <template v-else>
                {{ $filters.formatFloat(idea.score) }}
              </template>
            </th>

            <td>
              <a class="action-link link-danger" @click.prevent="showDeleteDialog(idea)">
                <MDBIcon icon="trash" iconStyle="fas"/>
              </a>
            </td>
          </tr>
          </tbody>
        </MDBTable>
      </MDBCol>
    </BContent>

    <BIdeaEditModal caption="Editing the idea" id="ideaEdit" v-model="editIdeaShow" :on-save="load"
                    :data="editIdeaData">
      <MDBInput label="Caption" required v-model="editIdeaData.caption" maxLength="255"/>
      <MDBTextarea label="Comment" class="mt-3" v-model="editIdeaData.comment"/>
      <MDBInput label="Link" class="mt-3" v-model="editIdeaData.link" maxLength="255"/>
      <MDBSwitch label="Ready for dev" wrapper-class="mt-3" v-model="editIdeaData.ready_for_dev"/>
      <MDBSwitch label="Done" wrapper-class="mt-3" v-model="editIdeaData.done"/>
      <MDBSwitch label="Make me the idea owner" wrapper-class="mt-3" v-model="editIdeaData.make_me_owner"
                 v-if="!editIdeaOwner"/>
    </BIdeaEditModal>

    <BIdeaEditModal caption="Editing the idea reach" id="reachEdit" v-model="editReachShow" :on-save="load"
                    :data="editReachData">
      <div class="alert alert-info show" v-if="options.reach_description">
        <MDBIcon icon="info-circle" iconStyle="fas" class="me-1"/>
        {{ options.reach_description }}
      </div>
      <MDBInput label="Reach" class="active" type="number" required v-model.number="editReachData.reach"/>
      <MDBTextarea label="Comment" class="mt-3" v-model="editReachData.reach_comment"/>
    </BIdeaEditModal>

    <BIdeaEditModal caption="Editing the idea goals" v-model="editGoalShow" :on-save="load"
                    :data="editGoalsData">
      <template v-for="goal in options.goals" :key="goal.id">
        <MDBInput :label="goal.caption" class="active" type="number" step="any"
                  v-model.number="editGoalsData.goals[goal.id]"/>
        <MDBTextarea label="Comment" class="mt-1 mb-3" rows="2" v-model="editGoalsData.goals_comments[goal.id]"/>
      </template>
    </BIdeaEditModal>

    <BIdeaEditModal caption="Editing the idea confident" v-model="editConfidentShow" :on-save="load"
                    :data="editConfidentData">
      <div class="form-outline">
        <select class="form-control active" style="border: #bdbdbd solid thin" id="confident_select"
                v-model="editConfidentData.confident">
          <option v-for="cl in options.confident_levels" :key="cl.id" :value="cl.id">
            {{ cl.caption }}&nbsp;({{ cl.weight }})
          </option>
        </select>
        <label class="form-label bg-white ps-1 pe-1" for="confident_select">Confidence</label>
      </div>
      <MDBTextarea label="Comment" class="mt-3" v-model="editConfidentData.confident_comment"/>
    </BIdeaEditModal>

    <BIdeaEditModal caption="Editing the idea effort" v-model="editEffortShow" :on-save="load"
                    :data="editEffortData">
      <div class="alert alert-info show" v-if="options.effort_description">
        <MDBIcon icon="info-circle" iconStyle="fas" class="me-1"/>
        {{ options.effort_description }}
      </div>
      <template v-for="team in options.teams" :key="team.id">
        <MDBInput :label="team.caption" class="active" type="number"
                  v-model.number="editEffortData.teams[team.id]"/>
        <MDBTextarea label="Comment" class="mt-1 mb-3" rows="2" v-model="editEffortData.teams_comments[team.id]"/>
      </template>
    </BIdeaEditModal>

    <BConfirmModal v-model="deleteConfirmShow" yes-btn="Yes" no-btn="Cancel" :callback="deleteIdea">
      <p class="font-weight-bold">Are you sure to delete the idea "{{ deleteIdeaCur.caption }}"?</p>
      <p>The operation cannot be undone!</p>
    </BConfirmModal>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
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
  MDBSwitch,
  MDBTable,
  MDBTextarea
} from "mdb-vue-ui-kit";
import api, {
  IdeasListIdeaGoalV1,
  IdeasListIdeaTeamV1,
  IdeasListIdeaV1,
  IdeasListUserV1,
  ProjectsOptionsConfidentV1,
  ProjectsOptionsGoalV1,
  ProjectsOptionsOptionsV1,
  ProjectsOptionsTeamV1
} from "@/api";
import BContent from "@/components/BContent.vue";
import BIdeaEditModal from "@/components/BIdeaEditModal.vue";
import filters from "@/filters/format";
import BUser from "@/components/BUser.vue";
import BConfirmModal from "@/components/BConfirmModal.vue";
import BComment from "@/components/BComment.vue";

export default defineComponent({
  name: 'ProjectIdeas',
  components: {
    BComment,
    BConfirmModal,
    BUser,
    BIdeaEditModal,
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
    MDBSwitch,
    MDBTable,
    MDBTextarea
  },
  computed: {
    ideasView(): ideaView[] {
      return this.ideas.map(idea => {
        const goals = idea.goals ? idea.goals.reduce((res, goal) => {
              res[goal.id] = goal
              return res
            }, {} as { [key: number]: IdeasListIdeaGoalV1 }) : {},

            teams = idea.teams ? idea.teams.reduce((res, team) => {
              res[team.id] = team
              return res
            }, {} as { [key: number]: IdeasListIdeaTeamV1 }) : {},

            impact = idea.goals ? this.options.goals.reduce((sum, goal) => {
              return goals[goal.id] ? sum + goals[goal.id].value * 10 / goal.divider : sum
            }, 0) : undefined,

            effort = idea.teams ? idea.teams.reduce((sum, team) => {
              return sum + team.capacity
            }, 0) : undefined

        let score = idea.reach !== undefined && impact !== undefined && idea.confident != undefined
            ? idea.reach * impact * this.confidentLevelsMap[idea.confident].weight : undefined
        if (score && effort)
          score /= effort

        return {
          id: idea.id,
          caption: idea.caption,
          comment: idea.comment,
          ready_for_dev: idea.ready_for_dev,
          done: idea.done,
          link: idea.link,
          reach: idea.reach,
          reach_comment: idea.reach_comment,
          confident: idea.confident,
          confident_comment: idea.confident_comment,
          goals: goals,
          teams: teams,
          score: score,
          owner: idea.owner,
          filled_effort: this.options?.teams.filter((team => {
            return teams[team.id] === undefined
          })).length === 0
        }
      }).sort((a, b) => {
        return a.done == b.done ?
            a.filled_effort == b.filled_effort
                ? (b.score || 0) - (a.score || 0)
                : (b.filled_effort ? 1 : 0) - (a.filled_effort ? 1 : 0)
            : (a.done ? 1 : 0) - (b.done ? 1 : 0)
      })
    },

    curIdeaId(): number | undefined {
      const match = /idea(\d+)/.exec(this.$route.hash)
      return match ? parseInt(match[1]) : undefined
    },

    confidentLevelsMap(): { [key: number]: ProjectsOptionsConfidentV1 } {
      return this.options && this.options.confident_levels ? this.options.confident_levels.reduce((res, cl) => {
        res[cl.id] = cl
        return res
      }, {} as { [key: number]: ProjectsOptionsConfidentV1 }) : {}
    },
  },
  data() {
    return {
      ideas: [] as IdeasListIdeaV1[],
      loading: true,
      error: '',

      addModal: false,
      options: {} as ProjectsOptionsOptionsV1,
      optionsLoading: false,
      optionsError: '',
      addData: {
        caption: ''
      },
      adding: false,
      addError: '',

      editIdeaShow: false,
      editIdeaOwner: false,
      editIdeaData: {
        id: 0,
        caption: '',
        comment: '',
        link: '',
        ready_for_dev: false,
        done: false,
        make_me_owner: false,
      },

      editReachShow: false,
      editReachData: {
        id: 0,
        reach: 0 as number | undefined,
        reach_comment: '' as string | undefined
      },

      editGoalShow: false,
      editGoalsData: {
        id: 0,
        goals: {} as { [key: number]: number },
        goals_comments: {} as { [key: number]: string },
      },

      editConfidentShow: false,
      editConfidentData: {
        id: 0,
        confident: undefined as number | undefined,
        confident_comment: '' as string | undefined
      },

      editEffortShow: false,
      editEffortData: {
        id: 0,
        teams: {} as { [key: number]: number },
        teams_comments: {} as { [key: number]: string },
      },

      deleteConfirmShow: false,
      deleteIdeaCur: {} as ideaView,
    }
  },
  watch: {
    curIdeaId() {
      this.scrollToIdea()
    },
    ideasView() {
      this.$nextTick(() => {
        this.scrollToIdea()
      })
    }
  },
  mounted() {
    this.loadOptions()
  },
  methods: {
    scrollToIdea() {
      const ideaEl = document.getElementById(`idea${this.curIdeaId}`)
      ideaEl && ideaEl.scrollIntoView()
    },

    load(id?: number) {
      this.loading = true
      this.error = ''

      api.IdeasListV1({project_id: this.$route.params['id'] as string}).then(ideas => {
        this.ideas = ideas
        if (id) {
          location.hash = `#idea${id}`
        }
        setTimeout(this.scrollToIdea, 100)
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.loading = false
      })
    },

    loadOptions() {
      this.optionsLoading = true
      this.optionsError = ""

      api.ProjectsOptionsV1({project_id: this.$route.params['id'] as string}).then(options => {
        this.options = options
        this.load()
      }).catch(err => {
        this.optionsError = err
      }).finally(() => {
        this.optionsLoading = false
      })
    },

    formatNumber(n: number, format: number): string {
      switch (format) {
        case 0:
          return filters.formatInt(n)
        case 1:
          return filters.formatFloat(n) + '%'
        case 2:
          return this.options.money_symbol + filters.formatFloat(n)
        default:
          return filters.formatFloat(n)
      }
    },

    filteredGoals(idea: ideaView): ProjectsOptionsGoalV1[] {
      return this.options.goals.filter(v => {
        return idea.goals[v.id]?.value || idea.goals[v.id]?.comment
      })
    },

    filteredTeams(idea: ideaView): ProjectsOptionsTeamV1[] {
      if (Object.keys(idea.teams).length === 0)
        return []

      return this.options.teams.filter(v => {
        return idea.teams[v.id]?.capacity !== 0 || idea.teams[v.id]?.comment
      })
    },

    addForm() {
      this.addModal = true
    },

    add() {
      this.adding = true
      this.addError = ''

      api.IdeasAddV1({
        project_id: this.$route.params['id'] as string,
        caption: this.addData.caption
      }).then(() => {
        this.addModal = false
        this.load()
      }).catch(err => {
        this.addError = err
      }).finally(() => {
        this.adding = false
      })
    },

    editIdea(idea: ideaView) {
      this.editIdeaData = {
        id: idea.id,
        caption: idea.caption,
        comment: idea.comment || '',
        link: idea.link || '',
        ready_for_dev: idea.ready_for_dev,
        done: idea.done,
        make_me_owner: false
      }
      this.editIdeaShow = true
      this.editIdeaOwner = this.$store.state.user.email == idea.owner.email
    },

    editReach(idea: ideaView) {
      this.editReachData = {
        id: idea.id,
        reach: idea.reach,
        reach_comment: idea.reach_comment
      }
      this.editReachShow = true
    },

    editGoals(idea: ideaView) {
      this.editGoalsData = {
        id: idea.id,
        goals: this.options.goals.reduce((res, goal) => {
          res[goal.id] = idea.goals[goal.id]?.value
          return res
        }, {} as { [key: number]: number }),
        goals_comments: this.options.goals.reduce((res, goal) => {
          res[goal.id] = idea.goals[goal.id]?.comment || ''
          return res
        }, {} as { [key: number]: string }),
      }
      this.editGoalShow = true
    },

    editConfident(idea: ideaView) {
      this.editConfidentData = {
        id: idea.id,
        confident: idea.confident,
        confident_comment: idea.confident_comment
      }
      this.editConfidentShow = true
    },

    editEffort(idea: ideaView) {
      this.editEffortData = {
        id: idea.id,
        teams: this.options.teams.reduce((res, team) => {
          res[team.id] = idea.teams[team.id]?.capacity
          return res
        }, {} as { [key: number]: number }),
        teams_comments: this.options.teams.reduce((res, team) => {
          res[team.id] = idea.teams[team.id]?.comment || ''
          return res
        }, {} as { [key: number]: string }),
      }
      this.editEffortShow = true
    },

    showDeleteDialog(idea: ideaView) {
      this.deleteIdeaCur = idea
      this.deleteConfirmShow = true
    },

    // eslint-disable-next-line
    deleteIdea(): Promise<any> {
      return api.IdeasDeleteV1({id: this.deleteIdeaCur.id}).then(v => {
        // eslint-disable-next-line
        return new Promise<any>((resolve) => {
          this.load()
          resolve(v)
        })
      })
    },

    formKeyPress(e: KeyboardEvent) {
      if (e.keyCode === 13 && (e.target as HTMLElement).nodeName != "TEXTAREA") {
        e.preventDefault();
        this.add()
      }
    }

  },
});

interface ideaView {
  id: number
  caption: string
  comment?: string
  ready_for_dev: boolean
  done: boolean,
  link?: string
  reach?: number
  reach_comment?: string
  confident?: number
  confident_comment?: string
  goals: { [key: number]: IdeasListIdeaGoalV1 }
  teams: { [key: number]: IdeasListIdeaTeamV1 }
  score?: number
  owner: IdeasListUserV1
}
</script>

<style lang="scss">
.ideas-tbl {
  overflow: hidden;
}

.edit-icon {
  visibility: hidden;
  color: #00000080;
  cursor: pointer;
  position: absolute;
  top: 8px;
  left: 2px;
}

th:hover a.edit-link .edit-icon, td:hover a.edit-link .edit-icon {
  visibility: visible;
}

.action-link {
  visibility: hidden;
  cursor: pointer;
}

tr:hover a.action-link, tr:hover a.action-link {
  visibility: visible;
}

tr.current > *, .table-striped > tbody > tr.current:nth-of-type(odd) > * {
  animation: target-fade 3s 1;
}

@keyframes target-fade {
  0% {
    --mdb-table-accent-bg: transparent;
    background-color: #f0d8ff;
  }
  50% {
    --mdb-table-accent-bg: transparent;
    background-color: #f0d8ff;
  }
  100% {
    --mdb-table-accent-bg: default;
    background-color: transparent;
  }
}
</style>