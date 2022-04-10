import React from "react";
import {Coordinate} from "../entities/Coordinate";
import './CoordinateItem.scss'

export const CoordinateItem: React.FC<Coordinate> = ({title, description, x, y}) => {
    return (
        <article className="coordinate-item">
            <section className="coordinate-item-text">
                <h3>{title}</h3>
                <p>{description}</p>
            </section>

            <p>({x}, {y})</p>
        </article>
    )
}