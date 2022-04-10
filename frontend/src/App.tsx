import React from 'react';
import './App.css';
import {Header} from "./components/Header";
import {MapFrame} from "./components/MapFrame";

function App() {
    return (
        <div className="App">
            <Header />
            <MapFrame coordinates={[
                { title: "CoordinateItem 1", description: "This is c-1", x: 28751.408239482, y: 28550.6200304516 },
                { title: "CoordinateItem 2", description: "This is c-2", x: 34745.823492592, y: 34745.2983742922 }
            ]} />
        </div>
    );
}

export default App;
