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
    EntTitleEdges,
    EntTitleEdgesFromJSON,
    EntTitleEdgesFromJSONTyped,
    EntTitleEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTitle
 */
export interface EntTitle {
    /**
     * TITLE holds the value of the "TITLE" field.
     * @type {string}
     * @memberof EntTitle
     */
    tITLE?: string;
    /**
     * 
     * @type {EntTitleEdges}
     * @memberof EntTitle
     */
    edges?: EntTitleEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTitle
     */
    id?: number;
}

export function EntTitleFromJSON(json: any): EntTitle {
    return EntTitleFromJSONTyped(json, false);
}

export function EntTitleFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTitle {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'tITLE': !exists(json, 'TITLE') ? undefined : json['TITLE'],
        'edges': !exists(json, 'edges') ? undefined : EntTitleEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntTitleToJSON(value?: EntTitle | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'TITLE': value.tITLE,
        'edges': EntTitleEdgesToJSON(value.edges),
        'id': value.id,
    };
}


