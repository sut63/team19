/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntInstructorInfoEdges,
    EntInstructorInfoEdgesFromJSON,
    EntInstructorInfoEdgesFromJSONTyped,
    EntInstructorInfoEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntInstructorInfo
 */
export interface EntInstructorInfo {
    /**
     * EMAIL holds the value of the "EMAIL" field.
     * @type {string}
     * @memberof EntInstructorInfo
     */
    eMAIL?: string;
    /**
     * NAME holds the value of the "NAME" field.
     * @type {string}
     * @memberof EntInstructorInfo
     */
    nAME?: string;
    /**
     * PASSWORD holds the value of the "PASSWORD" field.
     * @type {string}
     * @memberof EntInstructorInfo
     */
    pASSWORD?: string;
    /**
     * PHONENUMBER holds the value of the "PHONENUMBER" field.
     * @type {string}
     * @memberof EntInstructorInfo
     */
    pHONENUMBER?: string;
    /**
     * 
     * @type {EntInstructorInfoEdges}
     * @memberof EntInstructorInfo
     */
    edges?: EntInstructorInfoEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntInstructorInfo
     */
    id?: number;
}

export function EntInstructorInfoFromJSON(json: any): EntInstructorInfo {
    return EntInstructorInfoFromJSONTyped(json, false);
}

export function EntInstructorInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInstructorInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'eMAIL': !exists(json, 'EMAIL') ? undefined : json['EMAIL'],
        'nAME': !exists(json, 'NAME') ? undefined : json['NAME'],
        'pASSWORD': !exists(json, 'PASSWORD') ? undefined : json['PASSWORD'],
        'pHONENUMBER': !exists(json, 'PHONENUMBER') ? undefined : json['PHONENUMBER'],
        'edges': !exists(json, 'edges') ? undefined : EntInstructorInfoEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntInstructorInfoToJSON(value?: EntInstructorInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'EMAIL': value.eMAIL,
        'NAME': value.nAME,
        'PASSWORD': value.pASSWORD,
        'PHONENUMBER': value.pHONENUMBER,
        'edges': EntInstructorInfoEdgesToJSON(value.edges),
        'id': value.id,
    };
}


