
import HomePage from "..."


class App extends React.Component {
    render() {
        return (
            <Router>
                <Route path="/" component={HomePage} />
            </Router>
        );
    }
}

const Renderable = connectAndRender(masterReducer, <App />);
