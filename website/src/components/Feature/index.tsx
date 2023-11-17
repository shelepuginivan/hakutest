import { ComponentProps, ComponentType, FC } from 'react'

import styles from './styles.module.css'

export interface FeatureProps {
    title: string
    description: string
    Svg: ComponentType<ComponentProps<'svg'>>
}

const Feature: FC<FeatureProps> = ({ title, description, Svg }) => {
    return (
        <div className={styles.feature}>
            <div className={styles.featureText}>
                <h3 className={styles.featureTitle}>{title}</h3>
                <p>{description}</p>
            </div>
            <Svg className={styles.featureImage} />
        </div>
    )
}

export default Feature
