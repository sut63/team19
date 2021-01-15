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
    EntSubjectsOfferedEdges,
    EntSubjectsOfferedEdgesFromJSON,
    EntSubjectsOfferedEdgesFromJSONTyped,
    EntSubjectsOfferedEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSubjectsOffered
 */
export interface EntSubjectsOffered {
    /**
     * AMOUNT holds the value of the "AMOUNT" field.
     * @type {string}
     * @memberof EntSubjectsOffered
     */
    aMOUNT?: string;
    /**
     * STATUS holds the value of the "STATUS" field.
     * @type {string}
     * @memberof EntSubjectsOffered
     */
    sTATUS?: string;
    /**
     * 
     * @type {number}
     * @memberof EntSubjectsOffered
     */
    degreeId?: number;
    /**
     * 
     * @type {EntSubjectsOfferedEdges}
     * @memberof EntSubjectsOffered
     */
    edges?: EntSubjectsOfferedEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSubjectsOffered
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof EntSubjectsOffered
     */
    subjectId?: number;
}

export function EntSubjectsOfferedFromJSON(json: any): EntSubjectsOffered {
    return EntSubjectsOfferedFromJSONTyped(json, false);
}

export function EntSubjectsOfferedFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSubjectsOffered {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'aMOUNT': !exists(json, 'AMOUNT') ? undefined : json['AMOUNT'],
        'sTATUS': !exists(json, 'STATUS') ? undefined : json['STATUS'],
        'degreeId': !exists(json, 'degree_id') ? undefined : json['degree_id'],
        'edges': !exists(json, 'edges') ? undefined : EntSubjectsOfferedEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'subjectId': !exists(json, 'subject_id') ? undefined : json['subject_id'],
    };
}

export function EntSubjectsOfferedToJSON(value?: EntSubjectsOffered | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'AMOUNT': value.aMOUNT,
        'STATUS': value.sTATUS,
        'degree_id': value.degreeId,
        'edges': EntSubjectsOfferedEdgesToJSON(value.edges),
        'id': value.id,
        'subject_id': value.subjectId,
    };
}

