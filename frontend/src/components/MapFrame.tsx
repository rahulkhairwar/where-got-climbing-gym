import React from "react";
import {Coordinate} from "../entities/Coordinate";
import {CoordinateItem} from "./CoordinateItem";
import './MapFrame.scss'

type Props = {
    coordinates: Coordinate[]
}

export const MapFrame: React.FC<Props> = ({coordinates}) => {
    return (
        <div className="map-frame">
            <ul className="coordinate-list">
                {
                    coordinates.map((c, i) => (
                        <li key={i}>
                            <CoordinateItem
                                title={c.title}
                                description={c.description}
                                x={c.x}
                                y={c.y}
                            />
                        </li>
                    ))
                }
            </ul>
        </div>
    )
}