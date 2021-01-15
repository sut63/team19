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
    EntInstructorRoomEdges,
    EntInstructorRoomEdgesFromJSON,
    EntInstructorRoomEdgesFromJSONTyped,
    EntInstructorRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntInstructorRoom
 */
export interface EntInstructorRoom {
    /**
     * BUILDING holds the value of the "BUILDING" field.
     * @type {string}
     * @memberof EntInstructorRoom
     */
    bUILDING?: string;
    /**
     * ROOM holds the value of the "ROOM" field.
     * @type {string}
     * @memberof EntInstructorRoom
     */
    rOOM?: string;
    /**
     * 
     * @type {EntInstructorRoomEdges}
     * @memberof EntInstructorRoom
     */
    edges?: EntInstructorRoomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntInstructorRoom
     */
    id?: number;
}

export function EntInstructorRoomFromJSON(json: any): EntInstructorRoom {
    return EntInstructorRoomFromJSONTyped(json, false);
}

export function EntInstructorRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntInstructorRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bUILDING': !exists(json, 'BUILDING') ? undefined : json['BUILDING'],
        'rOOM': !exists(json, 'ROOM') ? undefined : json['ROOM'],
        'edges': !exists(json, 'edges') ? undefined : EntInstructorRoomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntInstructorRoomToJSON(value?: EntInstructorRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'BUILDING': value.bUILDING,
        'ROOM': value.rOOM,
        'edges': EntInstructorRoomEdgesToJSON(value.edges),
        'id': value.id,
    };
}


