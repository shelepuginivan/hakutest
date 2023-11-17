import useDocusaurusContext from '@docusaurus/useDocusaurusContext'
import Layout from '@theme/Layout'

import HomepageHeader from '@site/src/components/HomepageHeader'

const Home = () => {
    const { siteConfig } = useDocusaurusContext()

    return (
        <Layout
            title={siteConfig.tagline}
            description={`${siteConfig.title} - ${siteConfig.tagline}`}
        >
            <HomepageHeader />
        </Layout>
    )
}

export default Home
