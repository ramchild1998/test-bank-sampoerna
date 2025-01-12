import RekeningDetail from '../../components/RekeningDetail'
import { useRouter } from 'next/router'

const RekeningPage = () => {
    const router = useRouter()
    const { id } = router.query

    return (
        <div>
            <RekeningDetail id={id} />
        </div>
    )
}

export default RekeningPage
