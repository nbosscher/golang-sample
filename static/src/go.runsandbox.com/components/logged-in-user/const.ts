

export enum ActionTypes {
    SAMPLE = "COMPONENT.SAMPLE"
}

export interface State {}

export interface StateProps {
    username: string;
}

export interface DispatchProps {}

export type Props = StateProps & DispatchProps;