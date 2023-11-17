import Link from '@docusaurus/Link'
import useDocusaurusContext from '@docusaurus/useDocusaurusContext'

import HakutestLogo from '@site/static/img/hakutest-logo.svg'

import styles from './styles.module.css'

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
                        Getting started 🚀
                    </Link>
                    <Link
                        className="button button--secondary button--lg"
                        to="/docs/intro"
                    >
                        Documentation 📗
                    </Link>
                </nav>
            </div>
        </header>
    )
}

export default HomepageHeader
