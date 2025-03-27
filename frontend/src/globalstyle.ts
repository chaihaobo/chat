import {createGlobalStyle} from "styled-components";

export default createGlobalStyle`
    :root {
        font-family: system-ui, Avenir, Helvetica, Arial, sans-serif;
        line-height: 1.5;
        font-weight: 400;

        color-scheme: light dark;

        font-synthesis: none;
        text-rendering: optimizeLegibility;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    html,body {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }

`