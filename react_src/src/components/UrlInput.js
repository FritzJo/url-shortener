import * as React from "react";

class UrlInput extends React.Component {
    constructor(props) {
        super(props);
        this.state = {value: ""};
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        this.setState({value: event.target.value});
    }

    handleSubmit(event) {
        const resultDiv = document.getElementById("result-div");
        document.getElementById("text_label").classList.add("active");
        resultDiv.style.display = "block";
        var currentHostname = window.location.hostname;

        var apiurl = "http://" + currentHostname + ":8080/api/v1/short";
        // Get target url
        var params = document.getElementById("icon_prefix").value;
        var xhr = new XMLHttpRequest();

        xhr.onreadystatechange = function () {
            if (this.readyState !== 4) {
                return;
            }
            if (this.status === 200) {
                var data = this.responseText;
                var json = JSON.parse(data);
                document.getElementById("s-url").innerHTML = json.target;
                document.getElementById("t-url").innerHTML = json.source;
            }

        };
        xhr.open("POST", apiurl, true);
        xhr.send(params);
        event.preventDefault();
    }

    render() {
        return (
            <div className="row">
                <div className="col s2"></div>
                <div className="input-field col s7">
                    <i className="material-icons prefix">wb_cloudy</i>
                    <input id="icon_prefix" type="text" className="validate"/>
                    <label id="text_label" htmlFor="icon_prefix">Target URL</label>
                </div>
                <div className="col s2">
                    <a onClick={this.handleSubmit} className="waves-effect waves-light btn"><i
                        className="material-icons left">near_me</i>Shrink!</a>
                </div>
                <div className="col s2"></div>
            </div>
        );
    }
}

export default UrlInput