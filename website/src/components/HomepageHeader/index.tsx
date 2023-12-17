import Link from '@docusaurus/Link'
import useDocusaurusContext from '@docusaurus/useDocusaurusContext'

import HakutestLogo from '@site/static/img/hakutest-logo.svg'

import styles from './styles.module.css'
import Translate from '@docusaurus/Translate'

const HomepageHeader = () => {
    const { siteConfig } = useDocusaurusContext()

    return (
        <header className={styles.header}>
            <HakutestLogo className={styles.logoSvg} />
            <div className={styles.headerContent}>
                <h2 className={styles.tagline}>{siteConfig.tagline}</h2>
                <nav className={styles.headerButtons}>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/intro"
                    >
                        <Translate>Getting started 🚀</Translate>
                    </Link>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/intro"
                    >
                        <Translate>Documentation 📗</Translate>
                    </Link>
                </nav>
            </div>
        </header>
    )
}

export default HomepageHeader
