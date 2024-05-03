import React from 'react'
import Header from '../components/defaultPage/Header';
import Footer from '../components/defaultPage/Footer'
import useDemoMutation from './useMutation'
import useDemoQuery from './useQuery'

function DemoPage() {
    const { onSubmit } = useDemoMutation();
    const { data } = useDemoQuery();
    return (
        <>
            <Header></Header>
            <Footer></Footer>
        </>
    )
}

export default DemoPage