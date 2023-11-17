import useDocusaurusContext from '@docusaurus/useDocusaurusContext'
import Layout from '@theme/Layout'

import HomepageHeader from '@site/src/components/HomepageHeader'

const Home = () => {
    const { siteConfig } = useDocusaurusContext()
    return (
        <Layout
            title={`Hello from ${siteConfig.title}`}
            description="Description will go into a meta tag in <head />"
        >
            <HomepageHeader />
        </Layout>
    )
}

export default Home
