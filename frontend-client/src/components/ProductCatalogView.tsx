import axios from 'axios'
import { useRouter } from 'next/router'
import React from 'react'
import { ImSpinner11 } from 'react-icons/im'

import { BASE_URL_PRODUCT_API } from '../../utils/const'
import { currencyIDR } from '../../utils/helpers'
import ActionTableButton from './ActionTableButton'
import Button from './Button'
import CustomTable from './CustomTable'

import Template from './Template'
import WrapperCard from './WrapperCard'

interface ProductProp {
    description: string,
    id: number,
    price: number,
    stock: number,
    title: string,
    is_master: string,
    created_at: number,
}

const ProductCatalogView = () => {
    const router = useRouter();
    const [product, setProduct] = React.useState<ProductProp[] | []>([])
    const [isLoading, setIsLoading] = React.useState<boolean>(true)
    const [isSync, setIsSync] = React.useState(false)

    const handleDeleteProduct = async (id: number) => {
        try {
            const response = await axios.delete(BASE_URL_PRODUCT_API + `/${id}`)
            fetchDataProduct()
        } catch (err) {
            console.log(err)
        }
    }

    const handleSyncProduct = async () => {
        try {
            setIsSync(true)
            const response = await axios.get(BASE_URL_PRODUCT_API + '/sync')
            setIsSync(false)
            fetchDataProduct()
        } catch (err) {
            setIsSync(false)
            console.log(err)
        }
    }

    const fetchDataProduct = async () => {
        try {
            const { data: result } = await axios.get(BASE_URL_PRODUCT_API)
            const { data } = result
            if (data === null) {
                setProduct([])
                setIsLoading(false)
                return
            }
            const filtererData = data.map((item: any) => {
                return {
                    ...item,
                    price: currencyIDR(item.price),
                    is_master: String(item.is_master),
                    action: item.is_master ? <ActionTableButton disabled /> : <ActionTableButton onUpdate={() => router.push(`/product/update/${item.id}`)} onDelete={() => handleDeleteProduct(item.id)} />
                }
            })
            setProduct([...filtererData])
            setIsLoading(false)
        } catch (err) {
            setProduct([])
            setIsLoading(false)
            console.log(err)
        }
    }
    const fetcher = React.useRef(fetchDataProduct);

    const columns = [
        {
            name: 'Id',
            selector: (row: any) => row.id
        },
        {
            name: 'Title',
            selector: (row: any) => row.title
        },
        {
            name: 'Stock',
            selector: (row: any) => row.stock
        },
        {
            name: 'Price',
            selector: (row: any) => row.price
        },
        {
            name: 'Is Master',
            selector: (row: any) => row.is_master
        },
        {
            name: 'Action',
            selector: (row: any) => row.action
        }
    ]

    React.useEffect(() => {
        fetcher.current()
    }, [fetcher])

    return (
        <Template title="Product Catalog">
            <WrapperCard>
                {isLoading ? (<div>loading</div>) : (<CustomTable columns={columns} data={product} />)}
            </WrapperCard>
            <div className="w-40">
                <Button onClick={handleSyncProduct} className="bg-primary text-white py-2 rounded mt-2 flex justify-center items-center">
                    {isSync ? (<ImSpinner11 size={20} className='animate-spin' />) : "Sync"}
                </Button>
            </div>
        </Template>
    )
}

export default ProductCatalogView