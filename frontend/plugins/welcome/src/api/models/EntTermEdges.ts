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
 * @interface EntTermEdges
 */
export interface EntTermEdges {
    /**
     * SubjectsOffered holds the value of the SubjectsOffered edge.
     * @type {Array<EntSubjectsOffered>}
     * @memberof EntTermEdges
     */
    subjectsOffered?: Array<EntSubjectsOffered>;
}

export function EntTermEdgesFromJSON(json: any): EntTermEdges {
    return EntTermEdgesFromJSONTyped(json, false);
}

export function EntTermEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTermEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'subjectsOffered': !exists(json, 'subjectsOffered') ? undefined : ((json['subjectsOffered'] as Array<any>).map(EntSubjectsOfferedFromJSON)),
    };
}

export function EntTermEdgesToJSON(value?: EntTermEdges | null): any {
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

