import React from "react";
import MovieBanner from "./MovieBanner";
import MovieList from "./MovieList";

function Content(username) {
    return (
        <div>
            <div className="py-5" style={{background: 'rgb(102, 67, 111)'}}>
                <MovieBanner {...username} />
            </div>
            <MovieList {...username} />
        </div>
    )
}

export default Content