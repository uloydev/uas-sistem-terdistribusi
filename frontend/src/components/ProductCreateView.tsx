import React, { useId } from 'react'
import axios from 'axios'
import { useRouter } from 'next/router'

import Input from './Input'
import Template from './Template'
import Label from './Label'
import TextArea from './TextArea'
import WrapperCard from './WrapperCard'
import Button from './Button'
import buttonStyles from '../../styles/Button.module.css'
import { BASE_URL_PRODUCT_API } from '../../utils/const'

const ProductCreateView = () => {
    const router = useRouter()
    const id = useId()
    const [description, setDescription] = React.useState("")
    const [title, setTitle] = React.useState("")
    const [price, setPrice] = React.useState<number>(0)
    const [stock, setStock] = React.useState<number>(0)

    const requestCreateProduct = async () => {
        const payload = {
            description,
            price,
            stock,
            title
        }
        try {
            await axios.post(`${BASE_URL_PRODUCT_API}?master=supersecret`, payload)
            router.push("/")
        }
        catch (err) {
            alert("something wrong in here")
            console.log(err)
        }

    }
    const handleChangeDesc = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        setDescription(e.target.value)
    }

    const handleChangeTitle = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTitle(e.target.value)
    }

    const handleChangePrice = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPrice(e.target.value as any)
    }

    const handleChangeStock = (e: React.ChangeEvent<HTMLInputElement>) => {
        setStock(e.target.value as any)
    }

    return (
        <Template title='Product Create'>
            <section className='mt-6'>
                <WrapperCard>
                    <div >
                        <h4 className='mb-7'>Create Product</h4>
                        <div className='space-y-4'>
                            <div>
                                <Label id={`${id}-title`}>Product Title</Label>
                                <Input
                                    id={`${id}-title`}
                                    value={title}
                                    type="text"
                                    onChange={handleChangeTitle}
                                    placeholder='Product title' />
                            </div>
                            <div>
                                <Label id={`${id}-price`}>Product Price</Label>
                                <Input
                                    id={`${id}-price`}
                                    value={price}
                                    type="number"
                                    onChange={handleChangePrice}
                                    placeholder='Product Price' />
                            </div>
                            <div>
                                <Label id={`${id}-stock`}>Product Stock</Label>
                                <Input
                                    id={`${id}-stock`}
                                    value={stock}
                                    type="number"
                                    onChange={handleChangeStock}
                                    placeholder='Product Stock' />
                            </div>
                            <div>
                                <Label id={`${id}-description`}>Product Description</Label>
                                <TextArea
                                    id={`${id}-description`}
                                    value={description}
                                    onChange={handleChangeDesc}
                                />
                            </div>
                        </div>
                        <div className='my-7 w-32'>
                            <Button
                                className={`${buttonStyles['btn-lavender']} hover:-translate-y-1.5 duration-200 py-2 rounded`}
                                onClick={requestCreateProduct}
                            >
                                Save
                            </Button>
                        </div>
                    </div>
                </WrapperCard>
            </section>
        </Template>
    )
}

export default ProductCreateView