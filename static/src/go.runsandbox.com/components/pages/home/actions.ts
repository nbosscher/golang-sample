
import {ActionTypes} from "./const"

export interface SampleActionParam {
    sample: string;
}

export function SampleAction(param: SampleActionParam) {
    return {
        type: ActionTypes.SAMPLE,
        payload: param,
    }
}