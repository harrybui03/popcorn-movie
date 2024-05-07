import React from 'react'
import useDemoQuery from './hook/useQuery';

function Demo() {
    // const { onSubmit } = useDemoMutation();
    const { data } = useDemoQuery();
    return (
        <>
            <h1>Huongdz</h1>            
        </>
    )
}

export default Demo