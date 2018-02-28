
import {ActionTypes} from "./const"

interface State {
    sample: string;
}

interface Action {
    type: ActionTypes;

}

const defaultState: State = {
    sample: "",
}

function reducer(state: State, action: Action): State {
    switch(action.type) {
        default:
            return state;
    }
}

export default reducer;