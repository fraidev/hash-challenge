// package: discount
// file: discount.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class DiscountRequest extends jspb.Message { 
    getUserId(): string;
    setUserId(value: string): DiscountRequest;
    clearProductsList(): void;
    getProductsList(): Array<Product>;
    setProductsList(value: Array<Product>): DiscountRequest;
    addProducts(value?: Product, index?: number): Product;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DiscountRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DiscountRequest): DiscountRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DiscountRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DiscountRequest;
    static deserializeBinaryFromReader(message: DiscountRequest, reader: jspb.BinaryReader): DiscountRequest;
}

export namespace DiscountRequest {
    export type AsObject = {
        userId: string,
        productsList: Array<Product.AsObject>,
    }
}

export class Product extends jspb.Message { 
    getId(): string;
    setId(value: string): Product;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Product.AsObject;
    static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Product;
    static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
    export type AsObject = {
        id: string,
    }
}

export class DiscountReply extends jspb.Message { 
    clearProductsList(): void;
    getProductsList(): Array<ProductDiscount>;
    setProductsList(value: Array<ProductDiscount>): DiscountReply;
    addProducts(value?: ProductDiscount, index?: number): ProductDiscount;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DiscountReply.AsObject;
    static toObject(includeInstance: boolean, msg: DiscountReply): DiscountReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DiscountReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DiscountReply;
    static deserializeBinaryFromReader(message: DiscountReply, reader: jspb.BinaryReader): DiscountReply;
}

export namespace DiscountReply {
    export type AsObject = {
        productsList: Array<ProductDiscount.AsObject>,
    }
}

export class ProductDiscount extends jspb.Message { 
    getPercentage(): number;
    setPercentage(value: number): ProductDiscount;
    getValueInCents(): number;
    setValueInCents(value: number): ProductDiscount;
    getProductId(): string;
    setProductId(value: string): ProductDiscount;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProductDiscount.AsObject;
    static toObject(includeInstance: boolean, msg: ProductDiscount): ProductDiscount.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProductDiscount, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProductDiscount;
    static deserializeBinaryFromReader(message: ProductDiscount, reader: jspb.BinaryReader): ProductDiscount;
}

export namespace ProductDiscount {
    export type AsObject = {
        percentage: number,
        valueInCents: number,
        productId: string,
    }
}
