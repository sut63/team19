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
    EntCourse,
    EntCourseFromJSON,
    EntCourseFromJSONTyped,
    EntCourseToJSON,
    EntSubjectsOffered,
    EntSubjectsOfferedFromJSON,
    EntSubjectsOfferedFromJSONTyped,
    EntSubjectsOfferedToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDegreeEdges
 */
export interface EntDegreeEdges {
    /**
     * Degree holds the value of the degree edge.
     * @type {Array<EntCourse>}
     * @memberof EntDegreeEdges
     */
    degree?: Array<EntCourse>;
    /**
     * SubjectsOffered holds the value of the SubjectsOffered edge.
     * @type {Array<EntSubjectsOffered>}
     * @memberof EntDegreeEdges
     */
    subjectsOffered?: Array<EntSubjectsOffered>;
}

export function EntDegreeEdgesFromJSON(json: any): EntDegreeEdges {
    return EntDegreeEdgesFromJSONTyped(json, false);
}

export function EntDegreeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDegreeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'degree': !exists(json, 'degree') ? undefined : ((json['degree'] as Array<any>).map(EntCourseFromJSON)),
        'subjectsOffered': !exists(json, 'subjectsOffered') ? undefined : ((json['subjectsOffered'] as Array<any>).map(EntSubjectsOfferedFromJSON)),
    };
}

export function EntDegreeEdgesToJSON(value?: EntDegreeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'degree': value.degree === undefined ? undefined : ((value.degree as Array<any>).map(EntCourseToJSON)),
        'subjectsOffered': value.subjectsOffered === undefined ? undefined : ((value.subjectsOffered as Array<any>).map(EntSubjectsOfferedToJSON)),
    };
}


