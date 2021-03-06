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
    EntSubjectsOffered,
    EntSubjectsOfferedFromJSON,
    EntSubjectsOfferedFromJSONTyped,
    EntSubjectsOfferedToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntYearEdges
 */
export interface EntYearEdges {
    /**
     * SubjectsOffered holds the value of the SubjectsOffered edge.
     * @type {Array<EntSubjectsOffered>}
     * @memberof EntYearEdges
     */
    subjectsOffered?: Array<EntSubjectsOffered>;
}

export function EntYearEdgesFromJSON(json: any): EntYearEdges {
    return EntYearEdgesFromJSONTyped(json, false);
}

export function EntYearEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntYearEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'subjectsOffered': !exists(json, 'subjectsOffered') ? undefined : ((json['subjectsOffered'] as Array<any>).map(EntSubjectsOfferedFromJSON)),
    };
}

export function EntYearEdgesToJSON(value?: EntYearEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'subjectsOffered': value.subjectsOffered === undefined ? undefined : ((value.subjectsOffered as Array<any>).map(EntSubjectsOfferedToJSON)),
    };
}


