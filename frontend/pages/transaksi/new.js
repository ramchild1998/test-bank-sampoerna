import TransaksiForm from '../../components/TransaksiForm'
import { useRouter } from 'next/router'

const NewTransaksi = () => {
    const router = useRouter()
    const { rekeningId } = router.query

    return (
        <div>
            <h1>Tambah Transaksi Baru</h1>
            <TransaksiForm rekeningId={rekeningId} onSave={() => window.location.href = `/rekening/${rekeningId}`} />
        </div>
    )
}

export default NewTransaksi
