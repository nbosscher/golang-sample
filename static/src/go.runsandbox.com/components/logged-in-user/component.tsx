
import {Props, State} from "./const"

class LoggedInUser extends React.Component<Props, State> {
    render() {
        return (
            <div>
                {this.props.username}
            </div>
        )
    }
}

export default LoggedInUser