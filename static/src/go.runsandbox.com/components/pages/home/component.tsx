
import {Props, State} from "./const"
import LoggedInUser from "go.runsandbox.com/components/logged-in-user"

class HomePage extends React.Component<Props, State> {
    render() {
        return (
            <div className="home-page">
                <div className="home-page_my-div"></div>
                <LoggedInUser />
            </div>
        )
    }
}

export default HomePage