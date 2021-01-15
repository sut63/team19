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
    EntClassroomEdges,
    EntClassroomEdgesFromJSON,
    EntClassroomEdgesFromJSONTyped,
    EntClassroomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClassroom
 */
export interface EntClassroom {
    /**
     * ROOM holds the value of the "ROOM" field.
     * @type {string}
     * @memberof EntClassroom
     */
    rOOM?: string;
    /**
     * 
     * @type {EntClassroomEdges}
     * @memberof EntClassroom
     */
    edges?: EntClassroomEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntClassroom
     */
    id?: number;
}

export function EntClassroomFromJSON(json: any): EntClassroom {
    return EntClassroomFromJSONTyped(json, false);
}

export function EntClassroomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClassroom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'rOOM': !exists(json, 'ROOM') ? undefined : json['ROOM'],
        'edges': !exists(json, 'edges') ? undefined : EntClassroomEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntClassroomToJSON(value?: EntClassroom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ROOM': value.rOOM,
        'edges': EntClassroomEdgesToJSON(value.edges),
        'id': value.id,
    };
}


