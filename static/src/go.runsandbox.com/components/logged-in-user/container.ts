
import RenderComponent from "./component"
import {StateProps, DispatchProps} from "./const"

const mapStateToProps = (state: AppState): StateProps => {
}

const mapDispatchToProps = (dispatcher: Dispatcher<AppState>, getState: Getter<AppState>): DispatchProps => {
}

const Component = connect(RenderComponent, mapStateToProps, mapDispatchToProps)
export default Component