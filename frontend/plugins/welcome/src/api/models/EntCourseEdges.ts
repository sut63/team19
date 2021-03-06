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
    EntDegree,
    EntDegreeFromJSON,
    EntDegreeFromJSONTyped,
    EntDegreeToJSON,
    EntDepartment,
    EntDepartmentFromJSON,
    EntDepartmentFromJSONTyped,
    EntDepartmentToJSON,
    EntSubject,
    EntSubjectFromJSON,
    EntSubjectFromJSONTyped,
    EntSubjectToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCourseEdges
 */
export interface EntCourseEdges {
    /**
     * 
     * @type {EntDegree}
     * @memberof EntCourseEdges
     */
    degreeID?: EntDegree;
    /**
     * 
     * @type {EntDepartment}
     * @memberof EntCourseEdges
     */
    departmentID?: EntDepartment;
    /**
     * 
     * @type {EntSubject}
     * @memberof EntCourseEdges
     */
    subjectID?: EntSubject;
}

export function EntCourseEdgesFromJSON(json: any): EntCourseEdges {
    return EntCourseEdgesFromJSONTyped(json, false);
}

export function EntCourseEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCourseEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'degreeID': !exists(json, 'DegreeID') ? undefined : EntDegreeFromJSON(json['DegreeID']),
        'departmentID': !exists(json, 'DepartmentID') ? undefined : EntDepartmentFromJSON(json['DepartmentID']),
        'subjectID': !exists(json, 'SubjectID') ? undefined : EntSubjectFromJSON(json['SubjectID']),
    };
}

export function EntCourseEdgesToJSON(value?: EntCourseEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'degreeID': EntDegreeToJSON(value.degreeID),
        'departmentID': EntDepartmentToJSON(value.departmentID),
        'subjectID': EntSubjectToJSON(value.subjectID),
    };
}


