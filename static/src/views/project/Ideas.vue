<template>
  <BContent :loading="loading || optionsLoading" :error="error || optionsError">
    <MDBCol md="12">
      <MDBBtn outline="primary" rounded size="md" style="margin-bottom: 0.5em;"
              aria-controls="addModal" @click="addForm">
        <MDBIcon icon="plus" iconStyle="fas" class="me-1"/>
        Add an idea
      </MDBBtn>

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

      <MDBTable striped sm>
        <thead>
        <tr>
          <th>Idea</th>
          <th>Reach
            <MDBTooltip v-model="tooltips['reach']" class="ms-1" style="cursor: help" v-if="options.reach_description">
              <template #reference>
                <MDBIcon icon="info-circle" iconStyle="fas"/>
              </template>
              <template #tip>
                {{ options.reach_description }}
              </template>
            </MDBTooltip>
          </th>
          <th>Impact</th>
          <th>Confidence</th>
          <th>Effort
            <MDBTooltip v-model="tooltips['effort']" class="ms-1" style="cursor: help"
                        v-if="options.effort_description">
              <template #reference>
                <MDBIcon icon="info-circle" iconStyle="fas"/>
              </template>
              <template #tip>
                {{ options.effort_description }}
              </template>
            </MDBTooltip>
          </th>
          <th>Score</th>
        </tr>
        </thead>

        <tbody>
        <tr v-for="idea in ideasView" :key="idea.id" :class="{'table-warning': idea.score === undefined}">
          <th class="position-relative">
            {{ idea.caption }}
            <a class="edit-link" @click.prevent="editCaption(idea)">
              <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
            </a>
          </th>

          <td class="position-relative">
            <template v-if="idea.reach !== undefined">{{ formatNumber(idea.reach, options.reach_format) }}</template>
            <span class="text-muted" v-else>&mdash;</span>
            <MDBTooltip v-model="tooltips[`reach${idea.id}`]" class="ms-1" style="cursor: help"
                        v-if="idea.reach_comment">
              <template #reference>
                <MDBIcon icon="comment" iconStyle="far"/>
              </template>
              <template #tip>
                {{ idea.reach_comment }}
              </template>
            </MDBTooltip>

            <a class="edit-link" @click.prevent="editReach(idea)">
              <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
            </a>
          </td>

          <td class="position-relative text-nowrap">
            <ul class="list-unstyled mb-0" v-if="filteredGoals(idea).length">
              <li v-for="goal in filteredGoals(idea)" :key="goal.id">
                <MDBBadge color="success" class="ms-1" style="width: 3em">
                  {{ idea.goals[goal.id].value * 10 / goal.divider }}
                </MDBBadge>
                {{ goal.caption }}:&nbsp;{{ formatNumber(idea.goals[goal.id].value, goal.format) }}
                <MDBTooltip v-model="tooltips[`goal${goal.id}_${idea.id}`]" class="ms-1" style="cursor: help"
                            v-if="idea.goals[goal.id].comment">
                  <template #reference>
                    <MDBIcon icon="comment" iconStyle="far"/>
                  </template>
                  <template #tip>
                    {{ idea.goals[goal.id]?.comment }}
                  </template>
                </MDBTooltip>
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
            <MDBTooltip v-model="tooltips[`confident${idea.id}`]" class="ms-1" style="cursor: help"
                        v-if="idea.confident_comment">
              <template #reference>
                <MDBIcon icon="comment" iconStyle="far"/>
              </template>
              <template #tip>
                {{ idea.confident_comment }}
              </template>
            </MDBTooltip>

            <a class="edit-link" @click.prevent="editConfident(idea)">
              <MDBIcon icon="edit" iconStyle="far" class="edit-icon"/>
            </a>
          </td>

          <td class="position-relative text-nowrap">
            <ul class="list-unstyled mb-0" v-if="filteredTeams(idea).length">
              <li v-for="team in filteredTeams(idea)" :key="team.id">{{ team.caption }}:&nbsp;
                {{ idea.teams[team.id].capacity }}
                <MDBTooltip v-model="tooltips[`team${team.id}_${idea.id}`]" class="ms-1" style="cursor: help"
                            v-if="idea.teams[team.id]?.comment">
                  <template #reference>
                    <MDBIcon icon="comment" iconStyle="far"/>
                  </template>
                  <template #tip>
                    {{ idea.teams[team.id]?.comment }}
                  </template>
                </MDBTooltip>
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
        </tr>
        </tbody>
      </MDBTable>
    </MDBCol>
  </BContent>

  <BIdeaEditModal caption="Editing the idea caption" id="captionEdit" v-model="editCaptionShow" :on-save="load"
                  :data="editCaptionData">
    <MDBInput label="Caption" required v-model="editCaptionData.caption"/>
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
      <MDBInput :label="goal.caption" class="active" type="number" step="any" required
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
      <MDBInput :label="team.caption" class="active" type="number" required
                v-model.number="editEffortData.teams[team.id]"/>
      <MDBTextarea label="Comment" class="mt-1 mb-3" rows="2" v-model="editEffortData.teams_comments[team.id]"/>
    </template>
  </BIdeaEditModal>

</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {
  MDBBadge,
  MDBBtn,
  MDBCol,
  MDBIcon,
  MDBInput,
  MDBModal,
  MDBModalBody,
  MDBModalFooter,
  MDBModalHeader,
  MDBModalTitle,
  MDBTable,
  MDBTextarea,
  MDBTooltip
} from "mdb-vue-ui-kit";
import api, {
  IdeasListIdeaGoalV1,
  IdeasListIdeaTeamV1,
  IdeasListIdeaV1,
  IdeasOptionsConfidentV1,
  IdeasOptionsGoalV1,
  IdeasOptionsOptionsV1,
  IdeasOptionsTeamV1
} from "@/api";
import BContent from "@/components/BContent.vue";
import BIdeaEditModal from "@/components/BIdeaEditModal.vue";
import filters from "@/filters/format";

export default defineComponent({
  name: 'ProjectIdeas',
  components: {
    BIdeaEditModal,
    BContent,
    MDBBadge,
    MDBBtn,
    MDBCol,
    MDBIcon,
    MDBInput,
    MDBModal,
    MDBModalHeader,
    MDBModalTitle,
    MDBModalBody,
    MDBModalFooter,
    MDBTable,
    MDBTextarea,
    MDBTooltip
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
          reach: idea.reach,
          reach_comment: idea.reach_comment,
          confident: idea.confident,
          confident_comment: idea.confident_comment,
          goals: goals,
          teams: teams,
          score: score
        }
      }).sort((a, b) => {
        return (b.score || 0) - (a.score || 0)
      })
    },

    confidentLevelsMap(): { [key: number]: IdeasOptionsConfidentV1 } {
      return this.options && this.options.confident_levels ? this.options.confident_levels.reduce((res, cl) => {
        res[cl.id] = cl
        return res
      }, {} as { [key: number]: IdeasOptionsConfidentV1 }) : {}
    },
  },
  data() {
    return {
      ideas: [] as IdeasListIdeaV1[],
      loading: true,
      error: '',

      tooltips: {} as { [key: string]: boolean },

      addModal: false,
      options: {} as IdeasOptionsOptionsV1,
      optionsLoading: false,
      optionsError: '',
      addData: {
        caption: ''
      },
      adding: false,
      addError: '',

      editCaptionShow: false,
      editCaptionData: {
        id: 0,
        caption: ''
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

    }
  },
  mounted() {
    this.loadOptions()
  },
  methods: {
    load() {
      this.loading = true
      this.error = ''

      api.IdeasListV1({project_id: parseInt(this.$route.params['id'] as string)}).then(ideas => {
        this.ideas = ideas
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.loading = false
      })
    },

    loadOptions() {
      this.optionsLoading = true
      this.optionsError = ""

      api.IdeasOptionsV1({project_id: parseInt(this.$route.params['id'] as string)}).then(options => {
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

    filteredGoals(idea: ideaView): IdeasOptionsGoalV1[] {
      return this.options.goals.filter(v => {
        return idea.goals[v.id]?.value
      })
    },

    filteredTeams(idea: ideaView): IdeasOptionsTeamV1[] {
      return this.options.teams.filter(v => {
        return idea.teams[v.id]
      })
    },

    addForm() {
      this.addModal = true
    },

    add() {
      this.adding = true
      this.addError = ''

      api.IdeasAddV1({
        project_id: parseInt(this.$route.params['id'] as string),
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

    editCaption(idea: ideaView) {
      this.editCaptionData = {
        id: idea.id,
        caption: idea.caption,
      }
      this.editCaptionShow = true
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
  reach?: number
  reach_comment?: string
  confident?: number
  confident_comment?: string
  goals: { [key: number]: IdeasListIdeaGoalV1 }
  teams: { [key: number]: IdeasListIdeaTeamV1 }
  score?: number
}
</script>

<style lang="scss">
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
</style>