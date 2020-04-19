import React from "react";
import "../css/styles.css"

class Header extends React.Component {
    render() {
        return (
            <div>
                <div className="row">
                    <div className="">
                        <h2 className="center-align">
                            <a href="/">URL Shortener</a>
                        </h2>
                    </div>
                </div>
            </div>
        );
    }
}

export default Header