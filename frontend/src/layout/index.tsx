import {FC, ReactNode} from "react";
import {styled} from "styled-components";


interface LayoutProps {
    children: ReactNode
}

const Index: FC<LayoutProps> = ({children}) => {
    return (
        <div>
            {children}
        </div>
    )
}

export default styled(Index)`
    
`