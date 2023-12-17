import Feature, { FeatureProps } from '@site/src/components/Feature'

import styles from './styles.module.css'
import Translate, { translate } from '@docusaurus/Translate'

const HomepageFeatures = () => {
    const features: FeatureProps[] = [
        {
            title: translate({ message: 'Efficient' }),
            description: translate({
                message:
                    'Hakutest is written in Go, making it a high-performance system.',
            }),
            Svg: require('@site/static/img/icon-speed.svg').default,
        },
        {
            title: translate({ message: 'Secure' }),
            description: translate({
                message:
                    'Your data, such as tests and results, is stored locally on your machine.',
            }),
            Svg: require('@site/static/img/icon-secure.svg').default,
        },
        {
            title: translate({ message: 'Configurable' }),
            description: translate({
                message:
                    'Hakutest provides an extensive configuration option, allowing you to tailor the experience to your specific needs.',
            }),
            Svg: require('@site/static/img/icon-gear.svg').default,
        },
        {
            title: translate({ message: 'Cross-platform' }),
            description: translate({
                message:
                    'Prebuilt packages are available for Windows, Linux, and macOS.',
            }),
            Svg: require('@site/static/img/icon-os.svg').default,
        },
    ]

    return (
        <section className={styles.homepageFeatures}>
            <h2 className={styles.featureListTitle}>
                <Translate>Features</Translate>
            </h2>
            <div className={styles.featureList}>
                {features.map((feature) => (
                    <Feature {...feature} />
                ))}
            </div>
        </section>
    )
}

export default HomepageFeatures
