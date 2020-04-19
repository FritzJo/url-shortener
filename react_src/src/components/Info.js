import React from "react";
import UrlInput from "./UrlInput"
import "../css/styles.css"

class Info extends React.Component {
    render() {
        return (<div>
                <UrlInput/>
                <div id="result-div" className="row hidden-div">
                    <div className="col s3"></div>
                    <div className="col s6">
                        <table>
                            <thead>
                            <tr>
                                <th>Short URL</th>
                                <th>Target URL</th>
                                <th>Expiration date</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr>
                                <td id="s-url">{this.props.short}</td>
                                <td id="t-url">{this.props.target}</td>
                                <td id="date-url">{this.props.date}</td>
                            </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        );
    }
}

export default Info