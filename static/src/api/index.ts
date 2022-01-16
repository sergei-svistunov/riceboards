export class ApiError extends Error {
    private readonly _code: string
    private readonly _message: string
    // eslint-disable-next-line
    private readonly _data: any

    // eslint-disable-next-line
    constructor(code: string, message: string, data: any) {
        super(message);
        this._code = code
        this._message = message
        this._data = data
    }

    get code(): string {
        return this._code;
    }

    get message(): string {
        return this._message;
    }

    get data(): string {
        return this._data;
    }
}

export default class API {
    private static readonly url = "/api"

    // eslint-disable-next-line
    private static post(method: string, request: any): Promise<any> {
        return fetch(
            this.url + method,
            {
                method: 'post',
                headers: {
                    'Content-Type': 'application/json',
                    'X-API-Key': localStorage.getItem('token') || ''
                },
                body: JSON.stringify(request)
            }
        )
            .then(response => {
                // eslint-disable-next-line
                return new Promise<any>((resolve, reject) => {
                    switch (response.status) {
                        case 200:
                            resolve(response)
                            break
                        case 400:
                            response.json().then(err => {
                                reject(new ApiError(err['code'], err['message'], err['data']))
                            })
                            break
                        default:
                            response.text().then(text => {
                                reject(new Error(text || response.statusText))
                            })
                    }
                })
            })
            .then((response) => response.json())
    }

    // Sign in through Google
    public static AuthGoogleV1(request: AuthGoogleReqV1): Promise<AuthGoogleUserV1> {
        return this.post('/auth/google/v1', request)
    }

    // Returns current user
    public static AuthMeV1(request: AuthMeReqV1): Promise<AuthMeUserV1> {
        return this.post('/auth/me/v1', request)
    }

    // Add a new idea
    public static IdeasAddV1(request: IdeasAddReqV1): Promise<IdeasAddIdeaV1> {
        return this.post('/ideas/add/v1', request)
    }

    // Delete the idea
    public static IdeasDeleteV1(request: IdeasDeleteReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/ideas/delete/v1', request)
    }

    // Add a new idea
    public static IdeasEditV1(request: IdeasEditReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/ideas/edit/v1', request)
    }

    // Returns projects list
    public static IdeasListV1(request: IdeasListReqV1): Promise<IdeasListIdeaV1[]> {
        return this.post('/ideas/list/v1', request)
    }

    // Add a new project
    public static ProjectsAddV1(request: ProjectsAddReqV1): Promise<ProjectsAddProjectV1> {
        return this.post('/projects/add/v1', request)
    }

    // Add a new project confidence level
    public static ProjectsConfidenceAddV1(request: ProjectsConfidenceAddReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/confidence/add/v1', request)
    }

    // Delete the project confidence level
    public static ProjectsConfidenceDeleteV1(request: ProjectsConfidenceDeleteReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/confidence/delete/v1', request)
    }

    // Edit the project confidence level
    public static ProjectsConfidenceEditV1(request: ProjectsConfidenceEditReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/confidence/edit/v1', request)
    }

    // Edit the project
    public static ProjectsEditV1(request: ProjectsEditReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/edit/v1', request)
    }

    // Returns the project
    public static ProjectsGetV1(request: ProjectsGetReqV1): Promise<ProjectsGetProjectV1> {
        return this.post('/projects/get/v1', request)
    }

    // Add a new project goal
    public static ProjectsGoalsAddV1(request: ProjectsGoalsAddReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/goals/add/v1', request)
    }

    // Delete the project goal
    public static ProjectsGoalsDeleteV1(request: ProjectsGoalsDeleteReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/goals/delete/v1', request)
    }

    // Edit the project goal
    public static ProjectsGoalsEditV1(request: ProjectsGoalsEditReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/goals/edit/v1', request)
    }

    // Returns projects list
    public static ProjectsListV1(request: ProjectsListReqV1): Promise<ProjectsListProjectV1[]> {
        return this.post('/projects/list/v1', request)
    }

    // Returns the ideas options
    public static ProjectsOptionsV1(request: ProjectsOptionsReqV1): Promise<ProjectsOptionsOptionsV1> {
        return this.post('/projects/options/v1', request)
    }

    // Add a new project team
    public static ProjectsTeamsAddV1(request: ProjectsTeamsAddReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/teams/add/v1', request)
    }

    // Delete the project team
    public static ProjectsTeamsDeleteV1(request: ProjectsTeamsDeleteReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/teams/delete/v1', request)
    }

    // Edit the project team
    public static ProjectsTeamsEditV1(request: ProjectsTeamsEditReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/teams/edit/v1', request)
    }

    // Add a new project user
    public static ProjectsUsersAddV1(request: ProjectsUsersAddReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/users/add/v1', request)
    }

    // Delete the project user
    public static ProjectsUsersDeleteV1(request: ProjectsUsersDeleteReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/users/delete/v1', request)
    }
}

export interface AuthGoogleReqV1 {
    credential: string;
}

export interface AuthGoogleUserV1 {
    id: number;
    email: string;
    fullname: string;
    avatar_url: string;
    token: string;
    permissions: string[];
}

// eslint-disable-next-line
export interface AuthMeReqV1 {
}

export interface AuthMeUserV1 {
    id: number;
    fullname: string;
    avatar_url: string;
    email: string;
    permissions: string[];
}

export interface IdeasAddIdeaV1 {
    id: number;
}

export interface IdeasAddReqV1 {
    project_id: string;
    caption: string;
    comment?: string;
}

export interface IdeasDeleteReqV1 {
    id: number;
}

export interface IdeasEditReqV1 {
    id: number;
    caption?: string;
    comment?: string;
    link?: string;
    ready_for_dev?: boolean;
    make_me_owner?: boolean;
    reach?: number;
    reach_comment?: string;
    confident?: number;
    confident_comment?: string;
    goals?: Record<number, number>;
    goals_comments?: Record<number, string>;
    teams?: Record<number, number>;
    teams_comments?: Record<number, string>;
}

export interface IdeasListIdeaGoalV1 {
    id: number;
    value: number;
    comment?: string;
    owner: IdeasListUserV1;
}

export interface IdeasListIdeaTeamV1 {
    id: number;
    capacity: number;
    comment?: string;
    owner: IdeasListUserV1;
}

export interface IdeasListIdeaV1 {
    id: number;
    caption: string;
    comment?: string;
    ready_for_dev: boolean;
    reach?: number;
    reach_comment?: string;
    link?: string;
    confident?: number;
    confident_comment?: string;
    goals: IdeasListIdeaGoalV1[];
    teams: IdeasListIdeaTeamV1[];
    owner: IdeasListUserV1;
}

export interface IdeasListReqV1 {
    project_id: string;
}

export interface IdeasListUserV1 {
    fullname: string;
    email: string;
    avatar_url: string;
}

export interface ProjectsAddProjectV1 {
    id: number;
}

export interface ProjectsAddReqV1 {
    id: string;
    caption: string;
}

export interface ProjectsConfidenceAddReqV1 {
    project_id: string;
    caption: string;
    weight: number;
}

export interface ProjectsConfidenceDeleteReqV1 {
    project_id: string;
    id: number;
}

export interface ProjectsConfidenceEditReqV1 {
    project_id: string;
    id: number;
    caption: string;
    weight: number;
}

export interface ProjectsEditReqV1 {
    id: string;
    caption: string;
    reach_format: number;
    reach_description: string;
    effort_description: string;
    money_symbol: string;
}

export interface ProjectsGetProjectV1 {
    id: string;
    caption: string;
}

export interface ProjectsGetReqV1 {
    id: string;
}

export interface ProjectsGoalsAddReqV1 {
    project_id: string;
    caption: string;
    format: number;
    divider: number;
}

export interface ProjectsGoalsDeleteReqV1 {
    project_id: string;
    id: number;
}

export interface ProjectsGoalsEditReqV1 {
    project_id: string;
    id: number;
    caption: string;
    format: number;
    divider: number;
}

export interface ProjectsListProjectV1 {
    id: string;
    caption: string;
}

// eslint-disable-next-line
export interface ProjectsListReqV1 {
}

export interface ProjectsOptionsConfidentV1 {
    id: number;
    caption: string;
    weight: number;
}

export interface ProjectsOptionsGoalV1 {
    id: number;
    caption: string;
    format: number;
    divider: number;
}

export interface ProjectsOptionsOptionsV1 {
    caption: string;
    reach_format: number;
    reach_description: string;
    effort_description: string;
    money_symbol: string;
    confident_levels: ProjectsOptionsConfidentV1[];
    goals: ProjectsOptionsGoalV1[];
    teams: ProjectsOptionsTeamV1[];
    users: ProjectsOptionsUserV1[];
}

export interface ProjectsOptionsReqV1 {
    project_id: string;
    with_users?: boolean;
}

export interface ProjectsOptionsTeamV1 {
    id: number;
    caption: string;
}

export interface ProjectsOptionsUserV1 {
    fullname: string;
    email: string;
    avatar_url: string;
}

export interface ProjectsTeamsAddReqV1 {
    project_id: string;
    caption: string;
}

export interface ProjectsTeamsDeleteReqV1 {
    project_id: string;
    id: number;
}

export interface ProjectsTeamsEditReqV1 {
    project_id: string;
    id: number;
    caption: string;
}

export interface ProjectsUsersAddReqV1 {
    project_id: string;
    e_mail: string;
}

export interface ProjectsUsersDeleteReqV1 {
    project_id: string;
    e_mail: string;
}

// eslint-disable-next-line
export interface Struct_c7d8f8ae {
}