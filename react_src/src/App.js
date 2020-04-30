import React from "react";

import Footer from "./components/Footer.js";
import Header from "./components/Header";
import Info from "./components/Info";

function App() {
    return (
        <body className="site-background">
        <main>
            <Header/>
            <Info date="Hello World"/>
        </main>
        <Footer/>
        </body>
    );
}

export default App;
