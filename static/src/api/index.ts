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

    // Returns the ideas options
    public static IdeasOptionsV1(request: IdeasOptionsReqV1): Promise<IdeasOptionsOptionsV1> {
        return this.post('/ideas/options/v1', request)
    }

    // Add a new project
    public static ProjectsAddV1(request: ProjectsAddReqV1): Promise<ProjectsAddProjectV1> {
        return this.post('/projects/add/v1', request)
    }

    // Edit the project
    public static ProjectsEditV1(request: ProjectsEditReqV1): Promise<Struct_c7d8f8ae> {
        return this.post('/projects/edit/v1', request)
    }

    // Returns the project
    public static ProjectsGetV1(request: ProjectsGetReqV1): Promise<ProjectsGetProjectV1> {
        return this.post('/projects/get/v1', request)
    }

    // Returns projects list
    public static ProjectsListV1(request: ProjectsListReqV1): Promise<ProjectsListProjectV1[]> {
        return this.post('/projects/list/v1', request)
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
    project_id: number;
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
    project_id: number;
}

export interface IdeasListUserV1 {
    fullname: string;
    email: string;
    avatar_url: string;
}

export interface IdeasOptionsConfidentV1 {
    id: number;
    caption: string;
    weight: number;
}

export interface IdeasOptionsGoalV1 {
    id: number;
    caption: string;
    format: number;
    divider: number;
}

export interface IdeasOptionsOptionsV1 {
    caption: string;
    reach_format: number;
    reach_description: string;
    effort_description: string;
    money_symbol: string;
    confident_levels: IdeasOptionsConfidentV1[];
    goals: IdeasOptionsGoalV1[];
    teams: IdeasOptionsTeamV1[];
    users: IdeasOptionsUserV1[];
}

export interface IdeasOptionsReqV1 {
    project_id: number;
    with_users?: boolean;
}

export interface IdeasOptionsTeamV1 {
    id: number;
    caption: string;
}

export interface IdeasOptionsUserV1 {
    fullname: string;
    email: string;
    avatar_url: string;
}

export interface ProjectsAddProjectV1 {
    id: number;
    caption: string;
}

export interface ProjectsAddReqV1 {
    caption: string;
}

export interface ProjectsEditReqV1 {
    id: number;
    caption: string;
    reach_format: number;
    reach_description: string;
    effort_description: string;
    money_symbol: string;
}

export interface ProjectsGetProjectV1 {
    id: number;
    caption: string;
}

export interface ProjectsGetReqV1 {
    id: number;
}

export interface ProjectsListProjectV1 {
    id: number;
    caption: string;
}

// eslint-disable-next-line
export interface ProjectsListReqV1 {
}

// eslint-disable-next-line
export interface Struct_c7d8f8ae {
}